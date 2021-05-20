package main

/**
 * Apply quicksort to put the array elements between position lowIndex and position highIndex into increasing order.
 */
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(A []int, lowIndex int, highIndex int) {
	// The list has length one or zero.
	// Nothing needs to be done, so just return from the subroutine.
	if lowIndex >= highIndex {
		return
	}

	// Apply quickSortStep and get the new pivot position.
	pivotPosition := quickSortStep(A, lowIndex, highIndex)
	// Then apply quicksort to sort the items that precede the pivot and the items that follow it.
	quickSort(A, lowIndex, pivotPosition-1)
	quickSort(A, pivotPosition+1, highIndex)
}

/**
 * Apply quickSortStep to the list of items in locations lowIndex through highIndex in the array A.
 * The value returned by this routine is the final position of the pivot item in the array.
 */
func quickSortStep(A []int, lowIndex int, highIndex int) int {
	// lowIndex = 0, highIndex = 5, A = [10, 5, 89, 1, 20, 99], pivot = 10

	// When we begin, A[lowIndex] is an available space, since its value has been moved into the local variable, pivot.
	pivot := A[lowIndex]

	// The numbers highIndex and lowIndex mark the endpoints of a range of numbers that have not yet been tested.
	// Decrease highIndex and increase lowIndex until they become equal, moving numbers bigger than pivot
	// so that they lie above highIndex and moving umbers less than the pivot so that they lie below lowIndex.

	// lowIndex = 0, highIndex = 3, A = [10, 5, 89, 1, 20, 99] // for highIndex > lowIndex && A[highIndex] >= pivot { highIndex-- }
	// lowIndex = 1, highIndex = 3, A = [1, 5, 89, 1*, 20, 99] // A[lowIndex] = A[highIndex]; lowIndex++
	// lowIndex = 2, highIndex = 3, A = [1, 5, 89, 1*, 20, 99] // for lowIndex < highIndex && A[lowIndex] <= pivot { lowIndex++ }
	// lowIndex = 2, highIndex = 2, A = [1, 5, 89*, 89, 20, 99] // A[highIndex] = A[lowIndex]; highIndex--
	for highIndex > lowIndex {
		// Move highIndex down past numbers greater than pivot. These numbers do not have to be moved.
		for highIndex > lowIndex && A[highIndex] >= pivot {
			highIndex--
		}
		if highIndex == lowIndex {
			break
		}
		
		// The number A[highIndex] is less than pivot. Move it into the available space at A[lowIndex], 
		// leaving an available space at A[highIndex]!!!
		A[lowIndex] = A[highIndex]
		lowIndex++

		//---------------------------------------------// 
		
		// Move lowIndex up past numbers less than pivot. These numbers do not have to be moved.
		for lowIndex < highIndex && A[lowIndex] <= pivot {
			lowIndex++
		}
		if highIndex == lowIndex {
			break
		}
		
		// The number A[lowIndex] is greater than pivot. Move it into the available space at A[highIndex], 
		// leaving an available space at A[lowIndex]!!!
		A[highIndex] = A[lowIndex]
		highIndex--
	}
	
	// At this point, lowIndex has become equal to highIndex, and there is an available space at that position.
	// This position lies between numbers less than pivot and numbers greater than pivot.
	// Put pivot in this space and return its location.

	// lowIndex = 2, highIndex = 2, A = [1, 5, 10, 89, 20, 99] // A[lowIndex] = pivot
	A[lowIndex] = pivot
	return lowIndex
}
