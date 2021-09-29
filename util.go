package main

import (
	"encoding/csv"
	"io"
	"log"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func GetFunctionName(i interface{}) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer())
	fnNameSplit := strings.Split(fn.Name(), ".")
	fnShortName := fnNameSplit[len(fnNameSplit)-1]

	return fnShortName
}

func getRandomMapKey(m ResultsMap) int {
	for k := range m.data {
		return k
	}

	return 0
}

func writeCSV(writer io.Writer, csvRecords [][]string) {
	w := csv.NewWriter(writer)

	for _, row := range csvRecords {
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func getDurationsAVG(durations []time.Duration) float64 {
	sum := 0.0

	for _, duration := range durations {
		sum += duration.Seconds()
	}

	return sum / float64(len(durations))
}

type fakeWritter struct{}

func (fw fakeWritter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func voidLogger() *log.Logger {
	return log.New(fakeWritter{}, "", log.Default().Flags())
}
