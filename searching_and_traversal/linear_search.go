package searching_and_traversal

func LinearSearch(arr []int, key int) int {
	for index, item := range arr {
		if item == key {
			return index
		}
	}

	return -1
}
