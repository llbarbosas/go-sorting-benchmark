package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

type ResultsMap map[int]map[string][]time.Duration
type SortingFunc func([]int)

var (
	sortingFuncs = []SortingFunc{
		BubbleSort, SelectionSort, ShellSort,
		MergeSort, QuickSort, RadixSort,
	}
	seed int64 = 999
	inc        = 1000
	max        = 20000
	stp        = 1000
	rpt        = 10
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	aSet := GenerateAllA(seed, inc, max, stp)

	results := getResultsMap(aSet, sortingFuncs)
	csvRecords := resultsMapToCSV(results)
	writeCSV(os.Stdout, csvRecords)
}

func getResultsMap(aSet [][]int, sortingFuncs []SortingFunc) ResultsMap {
	results := make(ResultsMap)

	for _, a := range aSet {
		n := len(a)
		results[n] = make(map[string][]time.Duration, len(sortingFuncs))

		for _, fn := range sortingFuncs {
			fnName := getFunctionName(fn)

			results[n][fnName] = make([]time.Duration, rpt)

			log.Printf("%s(%d)\n", fnName, n)

			for i := 0; i < rpt; i++ {
				aCopy := make([]int, len(a))
				copy(aCopy, a)

				start := time.Now()
				fn(aCopy)
				elapsed := time.Since(start)

				log.Printf("\t⊢-- %f\n", elapsed.Seconds())

				results[n][fnName][i] = elapsed
			}

		}
	}

	return results
}

func resultsMapToCSV(results ResultsMap) [][]string {
	csv := make([][]string, len(results)+1)

	row := 0

	sortedNs := make([]int, 0, len(results))

	for n := range results {
		sortedNs = append(sortedNs, n)
	}

	sort.Ints(sortedNs)

	fnNames := make([]string, 0, len(results[0]))

	for _, n := range sortedNs {
		nResults := results[n]

		csv[row] = make([]string, len(nResults)+1)

		if row == 0 {
			csv[row][0] = "n"
			headerCol := 1

			for fnName := range nResults {
				csv[row][headerCol] = fnName
				headerCol++
				fnNames = append(fnNames, fnName)
			}

			row++
			csv[row] = make([]string, len(nResults)+1)
		}

		csv[row][0] = fmt.Sprint(n)

		col := 1

		for _, fnName := range fnNames {
			fnDurations := nResults[fnName]

			sum := 0.0

			for _, duration := range fnDurations {
				sum += duration.Seconds()
			}

			// len(fnDurations) == rpt
			csv[row][col] = fmt.Sprintf("%.6f", sum/float64(rpt))
			col++
		}

		row++
	}

	return csv
}

func writeCSV(writer io.Writer, csvRecors [][]string) {
	w := csv.NewWriter(writer)

	for _, row := range csvRecors {
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
