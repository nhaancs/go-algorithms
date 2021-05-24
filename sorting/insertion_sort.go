package sorting

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

func InsertionSortRecursive(arr []int) {
	insertionSortR(arr, len(arr))
}

/*
Input: arr = [99 8 4 89 1], length = 5
Call stack:
insertionSortR(arr, 5) // insert arr[4] into arr[:4]
	insertionSortR(arr, 4) // insert arr[3] into arr[:3]
		insertionSortR(arr, 3) // insert arr[2] into arr[:2]
			insertionSortR(arr, 2) // insert arr[1] into arr[:1]
				insertionSortR(arr, 1) // length = 1 => return, do nothing 
*/
func insertionSortR(arr []int, length int) {
    if length <= 1 {
        return
	}
    insertionSortR( arr, length-1 );
 
    // Insert last element at its correct position
    // in sorted array.
    temp := arr[length-1];
    j := length-2;
 
    /* Move elements of arr[0..i-1], that are
    greater than key, to one position ahead
    of their current position */
    for j >= 0 && arr[j] > temp {
        arr[j+1] = arr[j];
        j--;
    }
    arr[j+1] = temp;
}