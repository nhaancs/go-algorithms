package main

import "fmt"

func main() {
	arr := []int{99, 8, 4, 80, 88, 89, 1}
	fmt.Println(arr)
	// BubbleSort(arr)
	BubbleSortRecursive(arr)
	fmt.Println(arr)
}

