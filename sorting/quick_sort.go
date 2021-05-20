package main

/**
 * Apply quicksort to put the array elements between
 * position lo and position hi into increasing order.
 */
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(A []int, lo int, hi int) {
	if hi <= lo {
		// The list has length one or zero.  Nothing needs
		// to be done, so just return from the subroutine.
		return
	}

	// Apply quickSortStep and get the new pivot position.
	// Then apply quicksort to sort the items that
	// precede the pivot and the items that follow it.
	pivotPosition := quickSortStep(A, lo, hi)
	quickSort(A, lo, pivotPosition-1)
	quickSort(A, pivotPosition+1, hi)
}

/**
* Apply quickSortStep to the list of items in locations lo through hi
* in the array A. The value returned by this routine is the final
* position of the pivot item in the array.
 */
func quickSortStep(A []int, lo int, hi int) int {
	pivot := A[lo] // Get the pivot value.

	// The numbers hi and lo mark the endpoints of a range
	// of numbers that have not yet been tested.  Decrease hi
	// and increase lo until they become equal, moving numbers
	// bigger than pivot so that they lie above hi and moving
	// numbers less than the pivot so that they lie below lo.
	// When we begin, A[lo] is an available space, since its
	// value has been moved into the local variable, pivot.
	for hi > lo {
		// A[i] <= pivot
		// for i < lo, and A[i] >= pivot for i > hi.
		for hi > lo && A[hi] >= pivot {
			// Move hi down past numbers greater than pivot.
			// These numbers do not have to be moved.
			hi--
		}

		if hi == lo {
			break
		}

		// The number A[hi] is less than pivot.  Move it into
		// the available space at A[lo], leaving an available
		// space at A[hi].
		A[lo] = A[hi]
		lo++

		for hi > lo && A[lo] <= pivot {
			// Move lo up past numbers less than pivot.
			// These numbers do not have to be moved.
			lo++
		}

		if hi == lo {
			break
		}

		// The number A[lo] is greater than pivot.  Move it into
		// the available space at A[hi], leaving an available
		// space at A[lo].
		A[hi] = A[lo]
		hi--

	}

	// At this point, lo has become equal to hi, and there is
	// an available space at that position.  This position lies
	// between numbers less than pivot and numbers greater than
	// pivot.  Put pivot in this space and return its location.
	A[lo] = pivot
	return lo
}
