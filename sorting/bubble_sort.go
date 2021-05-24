package sorting

func BubbleSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for length > 1 {
		for i := 0; i < length-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		length--
	}
}

func BubbleSortRecursive(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length - 1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	BubbleSortRecursive(arr[:length-1])
}
