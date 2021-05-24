package main

import (
	"fmt"
	"go_data_structures_algorithms/sorting"
)

func main() {
	arr := []int{99, 8, 4, 80, 88, 89, 1, 3}
	fmt.Println(arr)
	// sorting.BubbleSort(arr)
	// sorting.BubbleSortRecursive(arr)
	// sorting.SelectionSort(arr)
	// sorting.SelectionSortRecursive(arr)
	// sorting.InsertionSort(arr)
	// sorting.InsertionSortRecursive(arr)
	// sorting.MergeSort(arr)
	sorting.QuickSort(arr)
	fmt.Println(arr)
}

