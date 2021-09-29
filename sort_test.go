package main_test

import (
	"reflect"
	"testing"

	main "github.com/llbarbosas/go-sorting-benchmark"
)

var (
	rand = main.MakeRandSource(999)
)

func BenchmarkBubbleSort(b *testing.B) {
	benchmarkSortingFunc(10, main.BubbleSort, b)
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmarkSortingFunc(10, main.SelectionSort, b)
}

func BenchmarkShellSort(b *testing.B) {
	benchmarkSortingFunc(10, main.ShellSort, b)
}

func BenchmarkMergeSort(b *testing.B) {
	benchmarkSortingFunc(10, main.MergeSort, b)
}

func BenchmarkQuickSort(b *testing.B) {
	benchmarkSortingFunc(10, main.QuickSort, b)
}

func BenchmarkRadixSort(b *testing.B) {
	benchmarkSortingFunc(10, main.RadixSort, b)
}

func TestBubbleSort(t *testing.T) {
	testSortingFunc(main.BubbleSort, t)
}

func TestSelectionSort(t *testing.T) {
	testSortingFunc(main.SelectionSort, t)
}

func TestShellSort(t *testing.T) {
	testSortingFunc(main.ShellSort, t)
}

func TestMergeSort(t *testing.T) {
	testSortingFunc(main.MergeSort, t)
}

func TestQuickSort(t *testing.T) {
	testSortingFunc(main.QuickSort, t)
}

func TestRadixSort(t *testing.T) {
	testSortingFunc(main.RadixSort, t)
}

func benchmarkSortingFunc(n int, fn main.SortingFunc, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := main.GenerateSingleA(n, rand)
		b.StartTimer()
		fn(s)
	}
}

func testSortingFunc(fn func([]int), t *testing.T) {
	arr := []int{4, 5, 2, 1, 3}
	expected := []int{1, 2, 3, 4, 5}

	fn(arr)

	if isEqual := reflect.DeepEqual(arr, expected); !isEqual {
		fnName := main.GetFunctionName(fn)

		t.Fatalf(
			"%s didn't order data sucessfully\nExpected: %v, got: %v",
			fnName, expected, arr,
		)
	}
}
