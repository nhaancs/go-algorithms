package arrays

// function to merge two sorted arrays into one array
// input1: [4, 6, 8] - increasing order
// input2: [3, 5, 7, 9] - increasing order
// output: [3, 4, 5, 6, 7, 8, 9] - increasing order // can be nil
func MergeSortedArrays(arr1, arr2 []int) []int {
	// check nil and empty arrays
	// get arr1 length
	arr1Total := len(arr1)
	if arr1Total == 0 {
		return arr2
	}
	// get arr2 length
	arr2Total := len(arr2)
	if arr2Total == 0 {
		return arr1
	}

	// create a new array to hold the result
	mergedArr := make([]int, arr1Total+arr2Total)
	arr1Item := arr1[0]
	arr2Item := arr2[0]
	next1 := 1
	next2 := 1
	for i := 0; i < len(mergedArr); i++ {
		if arr1Item < arr2Item {
			mergedArr[i] = arr1Item
			if next1 <= arr1Total-1 {
				arr1Item = arr1[next1]
				next1++
			}
		} else {
			mergedArr[i] = arr2Item
			if next2 <= arr2Total-1 {
				arr2Item = arr2[next2]
				next2++
			}
		}
	}

	return mergedArr
}
