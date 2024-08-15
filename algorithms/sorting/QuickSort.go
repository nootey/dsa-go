package sorting

// This implementation picks the last element as pivot

func partition(arr []int, start int, end int) int {
	// Pick pivot as the last item in the array
	pivot := arr[end]

	compareIndex := start
	smallestIndex := start - 1

	for ; compareIndex < end; compareIndex++ {
		// Check if value at compareIndex is less than pivot
		if arr[compareIndex] < pivot {
			smallestIndex++
			swapItem(arr, smallestIndex, compareIndex)
		}
	}

	// If pivot is smaller than unsorted value, swap them
	// The final resting place of the pivot, is the current smallest index

	// Increment the smallest index to put it into the correct position
	smallestIndex++

	swapItem(arr, smallestIndex, end)
	return smallestIndex
}

func swapItem(arr []int, i int, j int) {
	tempVar := arr[i]
	arr[i] = arr[j]
	arr[j] = tempVar
}

func QuickSort(arr []int, start int, end int) {

	// Base case
	if end <= start {
		return
	}

	pivot := partition(arr, start, end)
	// Sort the left side of pivot
	QuickSort(arr, start, pivot-1)
	// Sort the right side of pivot
	QuickSort(arr, pivot+1, end)
}
