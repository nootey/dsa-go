package sorting

import "dsa-go/structures/heap"

func HeapSort(data []int, isMinHeap bool) []int {
	h := heap.CreateNewHeap(isMinHeap)
	h.IsMinHeap = isMinHeap

	// Build the heap from the input data
	h.BuildHeap(data)

	// Extract elements from the heap and place them into the sorted portion
	sortedData := make([]int, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		sortedData[i] = h.ExtractRoot()
	}

	return sortedData
}
