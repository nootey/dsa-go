package heap

// This is an implementation of a binary heap

type Heap struct {
	Data      []int
	IsMinHeap bool
}

func CreateNewHeap() *Heap {
	return &Heap{
		Data: []int{},
	}
}

// Insert adds a new element to the heap and maintains the heap property
func (h *Heap) Insert(value int) {

	h.Data = append(h.Data, value)
	h.HeapifyUp(len(h.Data) - 1)
}

// HeapifyUp moves the newly added element up to maintain the heap property
func (h *Heap) HeapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2

		if h.compare(h.Data[index], h.Data[parentIndex]) {
			h.swap(index, parentIndex)
			index = parentIndex
		} else {
			break
		}
	}
}

// HeapifyDown moves the element at index down to maintain the heap property
func (h *Heap) HeapifyDown(index int) {
	lastIndex := len(h.Data) - 1
	l, r := 2*index+1, 2*index+2
	childToCompare := 0

	// Continue while there is at least one child
	for l <= lastIndex {
		if r <= lastIndex && h.compare(h.Data[r], h.Data[l]) { // Compare left and right children
			childToCompare = r
		} else {
			childToCompare = l
		}

		// Compare the selected child with the current node
		if h.compare(h.Data[childToCompare], h.Data[index]) {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = 2*index+1, 2*index+2
		} else {
			break
		}
	}
}

// compare returns true if the order is correct based on the heap type (min-heap or max-heap)
func (h *Heap) compare(child, parent int) bool {
	if h.IsMinHeap {
		return child < parent // min-heap: child should be less than parent
	}
	return child > parent // max-heap: child should be greater than parent
}

// swap swaps two elements in the array
func (h *Heap) swap(i1, i2 int) {
	h.Data[i1], h.Data[i2] = h.Data[i2], h.Data[i1]
}

// GetMin returns the minimum value for a min-heap, or the root in a max-heap
func (h *Heap) GetMin() int {
	if h.IsMinHeap && len(h.Data) > 0 {
		return h.Data[0]
	} else if !h.IsMinHeap && len(h.Data) > 0 {
		minVal := h.Data[len(h.Data)/2]
		for i := len(h.Data) / 2; i < len(h.Data); i++ {
			if h.Data[i] < minVal {
				minVal = h.Data[i]
			}
		}
		return minVal
	}
	return 0 // heap is empty
}

// GetMax returns the maximum value for a max-heap, or the root in a min-heap
func (h *Heap) GetMax() int {
	if !h.IsMinHeap && len(h.Data) > 0 {
		return h.Data[0]
	} else if h.IsMinHeap && len(h.Data) > 0 {
		maxVal := h.Data[len(h.Data)/2]
		for i := len(h.Data) / 2; i < len(h.Data); i++ {
			if h.Data[i] > maxVal {
				maxVal = h.Data[i]
			}
		}
		return maxVal
	}
	return 0 // heap is empty
}

// BuildHeap constructs a heap from an unsorted slice of integers
func (h *Heap) BuildHeap(data []int) {
	h.Data = data
	for i := len(data)/2 - 1; i >= 0; i-- {
		h.HeapifyDown(i)
	}
}

// ExtractRoot removes and returns the root of the heap
func (h *Heap) ExtractRoot() int {
	if len(h.Data) == 0 {
		return 0
	}

	// Swap the root with the last element
	root := h.Data[0]
	h.Data[0] = h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]

	h.HeapifyDown(0)

	return root
}
