package main_test

import (
	"reflect"
	"testing"

	main "github.com/llbarbosas/go-sorting-benchmark"
)

var rand = main.MakeRandSource(999)

func testSortingFunc(sortFunc func([]int)) bool {
	arr := []int{4, 5, 2, 1, 3}
	expected := []int{1, 2, 3, 4, 5}

	sortFunc(arr)

	return reflect.DeepEqual(arr, expected)
}

func TestBubbleSort(t *testing.T) {
	if !testSortingFunc(main.BubbleSort) {
		t.Fatal("BubbleSort not sorted sucessfully")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := main.GenerateSingleA(10, rand)
		b.StartTimer()
		main.BubbleSort(s)
	}
}

func TestSelectionSort(t *testing.T) {
	if !testSortingFunc(main.SelectionSort) {
		t.Fatal("SelectionSort not sorted sucessfully")
	}
}

func TestShellSort(t *testing.T) {
	if !testSortingFunc(main.ShellSort) {
		t.Fatal("ShellSort not sorted sucessfully")
	}
}

func TestMergeSort(t *testing.T) {
	if !testSortingFunc(main.MergeSort) {
		t.Fatal("MergeSort not sorted sucessfully")
	}
}

func TestQuickSort(t *testing.T) {
	if !testSortingFunc(main.QuickSort) {
		t.Fatal("QuickSort not sorted sucessfully")
	}
}

func TestRadixSort(t *testing.T) {
	if !testSortingFunc(main.RadixSort) {
		t.Fatal("RadixSort not sorted sucessfully")
	}
}
