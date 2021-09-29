package main

import (
	"log"
	"os"
)

var (
	aGenerationConfig = AGenerationConfig{
		Seed: 999,
		Inc:  1000,
		Max:  20000,
		Stp:  1000,
	}

	resultsMapGenerationConfig = ResultsMapGenerationConfig{
		SortingFuncs: []SortingFunc{
			BubbleSort, SelectionSort, ShellSort,
			MergeSort, QuickSort, RadixSort,
		},
		ASet:       GenerateASet(aGenerationConfig),
		Rpt:        10,
		DoParallel: true,
	}
)

func main() {
	results := NewResultsMap(resultsMapGenerationConfig)

	logger := log.New(os.Stderr, "", log.Ltime)
	results.Populate(logger)

	csvRecords := results.ToCSV()
	writeCSV(os.Stdout, csvRecords)
}
