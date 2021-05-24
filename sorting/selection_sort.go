package sorting

func SelectionSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i+1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func SelectionSortRecursive(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}

	minIndex := 0
	for i := 0; i < length; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}
	arr[0], arr[minIndex] = arr[minIndex], arr[0]
	SelectionSortRecursive(arr[1:])
}

