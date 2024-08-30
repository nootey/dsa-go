package searching

func BinarySearch(arr []int, target int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		middle := low + (high-low)/2
		value := arr[middle]

		if value == target {
			return middle
		} else if value > target {
			high = middle - 1
		} else if value < target {
			low = middle + 1
		}

	}

	return -1
}
