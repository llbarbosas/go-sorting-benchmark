package main

func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func SelectionSort(a []int) {
	length := len(a)

	for i := 0; i < length; i++ {
		maxIndex := 0

		for j := 1; j < length-i; j++ {
			if a[j] > a[maxIndex] {
				maxIndex = j
			}
		}

		a[length-i-1], a[maxIndex] = a[maxIndex], a[length-i-1]
	}
}

func ShellSort(a []int) {
	for d := int(len(a) / 2); d > 0; d /= 2 {
		for i := d; i < len(a); i++ {
			for j := i; j >= d && a[j-d] > a[j]; j -= d {
				a[j], a[j-d] = a[j-d], a[j]
			}
		}
	}
}

func MergeSort(a []int) {
	if len(a) <= 1 {
		return
	}

	middle := len(a) / 2
	MergeSort(a[:middle])
	MergeSort(a[middle:])

	aux := merge(a[:middle], a[middle:])

	copy(a, aux)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}

	return result
}

func QuickSort(a []int) {
	quickSortRecursion(a, 0, len(a)-1)
}

func quickSortRecursion(a []int, left int, right int) {
	if left < right {
		pivot := quickSortPartition(a, left, right)

		quickSortRecursion(a, left, pivot-1)
		quickSortRecursion(a, pivot+1, right)
	}
}

func quickSortPartition(a []int, left int, right int) int {
	for left < right {
		for left < right && a[left] <= a[right] {
			right--
		}

		if left < right {
			a[left], a[right] = a[right], a[left]
			left++
		}

		for left < right && a[left] <= a[right] {
			left++
		}

		if left < right {
			a[left], a[right] = a[right], a[left]
			right--
		}
	}
	return left
}

func RadixSort(a []int) {
	largestNum := radixSortGetBiggerNum(a)
	size := len(a)
	significantDigit := 1
	semiSorted := make([]int, size)

	for largestNum/significantDigit > 0 {
		bucket := [10]int{0}

		for i := 0; i < size; i++ {
			bucket[(a[i]/significantDigit)%10]++
		}

		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		for i := size - 1; i >= 0; i-- {
			bucket[(a[i]/significantDigit)%10]--
			semiSorted[bucket[(a[i]/significantDigit)%10]] = a[i]
		}

		for i := 0; i < size; i++ {
			a[i] = semiSorted[i]
		}

		significantDigit *= 10
	}
}

func radixSortGetBiggerNum(a []int) int {
	largestNum := 0

	for i := 0; i < len(a); i++ {
		if a[i] > largestNum {
			largestNum = a[i]
		}
	}

	return largestNum
}
