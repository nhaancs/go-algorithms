package main

import "big_o/demos"

func main() {
	arr := make([]string, 1000)
	arr[999] = "nemo"
	boxes := []string{"a", "b", "c", "d", "e"}

	// O(n) 
	demos.FindNemo(arr)

	// O(1)
	demos.LogFirstTwoElements(arr)

	// O(n^2)
	demos.LogAllPairsOfArr(boxes)
}
