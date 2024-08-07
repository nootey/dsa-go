package sorting

func InsertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		sortedIndex := i - 1

		for ; sortedIndex >= 0 && arr[sortedIndex] > key; sortedIndex-- {
			arr[sortedIndex+1] = arr[sortedIndex]
		}
		arr[sortedIndex+1] = key
	}
}
