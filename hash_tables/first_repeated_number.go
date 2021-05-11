package main

// Funtion to find the first repeated number in an given array
// Input: array of number (positive)
// Output: the numeber get repeated the first time, or -1 if no number gets repeated]
// Input = [2, 5, 1, 2, 3, 5, 1, 2, 4] => Output = 2
// Input = [2, 1, 1, 2, 3, 3, 1, 2, 4] => Output = 1
// Input = [2, 3, 4, 5] => Output = -1
func FirstRepeatedNumber(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	// naive solution: nested loops => inefficient

	// create a map has key is an int, and value is an bool
	elementsMap := map[int]bool{}
	// loop through the array
	for _, item := range arr {
		// check if current element is in a map
		if elementsMap[item] {// if yes, return this element
			return item
		}
		// if no, add this element to a map as a key, set the value to true
		elementsMap[item] = true
	}
		
	return -1
}