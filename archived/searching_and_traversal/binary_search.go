package searching_and_traversal

import "math"

// or lookup method of a binary search tree

func BinarySearch(sortedArr []int, key int) int {
	if len(sortedArr) < 1 {
		return -1
	}

	mid := int(math.Floor(float64(len(sortedArr)) / 2))
	if key == sortedArr[mid] {
		return mid
	}

	if key < sortedArr[mid] {
		return BinarySearch(sortedArr[:mid], key)
	}

	if key > sortedArr[mid] {
		return BinarySearch(sortedArr[mid+1:], key)
	}

	return -1
}
