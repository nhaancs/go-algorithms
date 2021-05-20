package main

import "fmt"

func main() {
	arr := []int{99, 8, 4, 80, 88, 89, 1, 3}
	fmt.Println(arr)
	// BubbleSort(arr)
	// BubbleSortRecursive(arr)
	// SelectionSort(arr)
	// SelectionSortRecursive(arr)
	// InsertionSort(arr)
	// InsertionSortRecursive(arr)
	// MergeSort(arr)
	QuickSort(arr)
	fmt.Println(arr)
}

