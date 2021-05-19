package main

func InsertionSort(arr []int) {
	for itemSorted := 0; itemSorted < len(arr); itemSorted++ {
		// Assume that items A[0], A[1], ... A[itemsSorted-1] have already been sorted. 
		// Insert A[itemsSorted] into the sorted part of the list
		temp := arr[itemSorted] // The item to be inserted
		loc := itemSorted-1 // Start at end of list
		for loc >= 0 && arr[loc] > temp {
			arr[loc+1] = arr[loc] // Bump item from A[loc] up to loc+1
			loc-- // Go on to next location
		}

		arr[loc+1] = temp // Put temp in last vacated space
	}
}
