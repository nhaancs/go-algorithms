package main

import (
	"math"
)

func MergeSort(arr []int) {
	// todo: why `arr = mergeSort(arr)` not working?
	// https://blog.golang.org/slices-intro
	arr = append(arr[:0], mergeSort(arr)...)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	// split array into right and left
	middle := int(math.Floor(float64(length)/2))
	left := arr[:middle]
	right := arr[middle:]
	// fmt.Println("middle", middle, "left: ", left, " - ", "right: ", right)

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := []int{}
	leftIndex := 0
	rightIndex := 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result =  append(result, right[rightIndex:]...)
	return result
}
