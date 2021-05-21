package main

func LinearSearch(arr []int, key int) int {
	for index, item := range arr {
		if item == key {
			return index
		}
	}

	return -1
}
