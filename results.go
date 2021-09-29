package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"time"
)

type ResultsMap struct {
	data map[int]map[string][]time.Duration
	mu   sync.Mutex
	cfg  ResultsMapGenerationConfig
}

func (rm ResultsMap) add(n int, fnName string, i int, duration time.Duration) {
	rm.mu.Lock()
	rm.data[n][fnName][i] = duration
	rm.mu.Unlock()
}

type ResultsMapGenerationConfig struct {
	Rpt          int
	ASet         [][]int
	SortingFuncs []SortingFunc
	DoParallel   bool
}

func (cfg ResultsMapGenerationConfig) SortingFuncsAsMap() map[string]SortingFunc {
	fnMap := make(map[string]SortingFunc)

	for _, fn := range cfg.SortingFuncs {
		fnName := GetFunctionName(fn)
		fnMap[fnName] = fn
	}

	return fnMap
}

func NewResultsMap(cfg ResultsMapGenerationConfig) ResultsMap {
	data := make(map[int]map[string][]time.Duration, len(cfg.ASet))

	for _, a := range cfg.ASet {
		n := len(a)

		data[n] = make(map[string][]time.Duration, n)

		for fnName := range cfg.SortingFuncsAsMap() {
			data[n][fnName] = make([]time.Duration, cfg.Rpt)
		}
	}

	return ResultsMap{
		data: data,
		cfg:  cfg,
	}
}

func (rm ResultsMap) Populate(customLogger ...*log.Logger) {
	var (
		wg     sync.WaitGroup
		logger *log.Logger
	)

	if len(customLogger) > 0 {
		logger = customLogger[0]
	} else {
		logger = voidLogger()
	}

	for _, a := range rm.cfg.ASet {
		n := len(a)

		for fnName, fn := range rm.cfg.SortingFuncsAsMap() {
			for i := 0; i < rm.cfg.Rpt; i++ {
				aCopy := make([]int, len(a))
				copy(aCopy, a)

				wg.Add(1)

				execFunc := func(fn SortingFunc, i int, ac []int) {
					start := time.Now()
					fn(ac)
					elapsed := time.Since(start)

					logger.Printf("[%d/%d] %s(%d): %f\n", i+1, rm.cfg.Rpt, fnName, n, elapsed.Seconds())

					rm.add(n, fnName, i, elapsed)

					wg.Done()
				}

				if rm.cfg.DoParallel {
					go execFunc(fn, i, aCopy)
				} else {
					execFunc(fn, i, aCopy)
				}
			}

			wg.Wait()
		}
	}
}

func (rm ResultsMap) ToCSV() [][]string {
	csv := make([][]string, len(rm.data)+1)

	orderedNs := rm.Ns()
	fnNames := rm.FnNames()

	csv[0] = append([]string{"n"}, fnNames...)

	row := 1

	for _, n := range orderedNs {
		nResults := rm.data[n]

		csv[row] = make([]string, len(nResults)+1)

		csv[row][0] = fmt.Sprint(n)

		col := 1

		for _, fnName := range fnNames {
			fnDurations := nResults[fnName]
			durationsAVG := getDurationsAVG(fnDurations)

			csv[row][col] = fmt.Sprintf("%.6f", durationsAVG)
			col++
		}

		row++
	}

	return csv
}

func (rm ResultsMap) FnNames() []string {
	fnNames := make([]string, 0, len(rm.cfg.SortingFuncs))

	for fnName := range rm.cfg.SortingFuncsAsMap() {
		fnNames = append(fnNames, fnName)
	}

	return fnNames
}

func (rm ResultsMap) Ns() []int {
	orderedNs := make([]int, 0, len(rm.data))

	for n := range rm.data {
		orderedNs = append(orderedNs, n)
	}

	sort.Ints(orderedNs)

	return orderedNs
}
