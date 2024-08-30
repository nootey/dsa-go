package searching

func LinearSearch(arr []int, value int) int {

	for i, v := range arr {
		if value == v {
			return i
		}
	}
	return -1
}
