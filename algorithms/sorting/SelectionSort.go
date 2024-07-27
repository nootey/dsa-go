package sorting

func SelectionSort(arr []int) {
	minEl := 0

	for i := 0; i < len(arr)-1; i++ {
		minEl = i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minEl] {
				minEl = j
			}
		}

		if minEl != i {
			temp := arr[i]
			arr[i] = arr[minEl]
			arr[minEl] = temp
		}
	}
}
