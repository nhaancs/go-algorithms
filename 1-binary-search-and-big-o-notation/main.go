package main

import "log"

func main() {
	log.Println(binarySearch([]int{}, 1))
	log.Println(binarySearch([]int{1, 2, 3}, 1))
	log.Println(binarySearch([]int{1, 2, 3}, 2))
}

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == item {
			return mid
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
