package main

import (
	"math/rand"
	"reflect"
	"runtime"
	"strings"
)

func MakeRandSource(seed int64) rand.Rand {
	source := rand.NewSource(seed)
	return *rand.New(source)
}

func GenerateAllA(seed int64, inc, max, stp int) [][]int {
	rand := MakeRandSource(seed)

	aNum := ((max - inc) / stp) + 1
	aSet := make([][]int, aNum)

	for n, i := inc, 0; n <= max; n, i = n+stp, i+1 {
		aSet[i] = GenerateSingleA(n, rand)
	}

	return aSet
}

func GenerateSingleA(n int, rand rand.Rand) []int {
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = rand.Int()
	}

	return a
}

func getFunctionName(i interface{}) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer())
	fnNameSplit := strings.Split(fn.Name(), ".")
	fnShortName := fnNameSplit[len(fnNameSplit)-1]

	return fnShortName
}
