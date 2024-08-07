package sorting

func MergeSort(arr []int, left int, right int) {

	// a subarray with one element is inherently sorted
	if left >= right {
		return
	}

	// find the middle index of the subarray
	mid := left + (right-left)/2

	// sort left side
	MergeSort(arr, left, mid)
	// sort right side
	MergeSort(arr, mid+1, right)

	// sort function
	merge(arr, left, mid, right)

}

func merge(arr []int, left int, mid int, right int) {

	// Create temporary slices
	leftArr := make([]int, mid-left+1)
	rightArr := make([]int, right-mid)

	// Copy values into the slices
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])

	fIndex := 0    // Initial index of first subarray
	sIndex := 0    // Initial index of second subarray
	mIndex := left // Initial index of merged subarray

	// Merge the two temporary arrays
	for fIndex < len(leftArr) && sIndex < len(rightArr) {
		if leftArr[fIndex] <= rightArr[sIndex] {
			arr[mIndex] = leftArr[fIndex]
			fIndex++
		} else {
			arr[mIndex] = rightArr[sIndex]
			sIndex++
		}
		mIndex++
	}

	// Copy any remaining elements to the main array
	for fIndex < len(leftArr) {
		arr[mIndex] = leftArr[fIndex]
		fIndex++
		mIndex++
	}

	for sIndex < len(rightArr) {
		arr[mIndex] = rightArr[sIndex]
		sIndex++
		mIndex++
	}

}
