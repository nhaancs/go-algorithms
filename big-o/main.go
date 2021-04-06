package main

func main() {
	arr := make([]string, 1000)
	arr[999] = "nemo"
	boxes := []string{"a", "b", "c", "d", "e"}

	// O(n) 
	findNemo(arr)

	// O(1)
	logFirstTwoElements(arr)

	// O(n^2)
	logAllPairsOfArr(boxes)
}
