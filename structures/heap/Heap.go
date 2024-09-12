package heap

type Heap struct {
	data      []int
	isMinHeap bool
}

// Insert adds a new element to the heap and maintains the heap property
func (h *Heap) Insert(value int) {
	// TODO: Implement logic
}

// Peek returns the top element (min or max depending on heap type) without removing it
func (h *Heap) Peek() int {
	// TODO: Implement logic
	return 0
}

// Extract removes and returns the top element (min or max) from the heap
func (h *Heap) Extract() int {
	// TODO: Implement logic
	return 0
}

// HeapifyUp moves the newly added element up to maintain the heap property
func (h *Heap) HeapifyUp(index int) {
	// TODO: Implement logic
}

// HeapifyDown moves the element at index down to maintain the heap property
func (h *Heap) HeapifyDown(index int) {
	// TODO: Implement logic
}

// BuildHeap constructs a heap from an unsorted slice of integers
func (h *Heap) BuildHeap(data []int) {
	// TODO: Implement logic
}

// Size returns the number of elements in the heap
func (h *Heap) Size() int {
	// TODO: Implement logic
	return 0
}

// IsEmpty returns true if the heap is empty, false otherwise
func (h *Heap) IsEmpty() bool {
	// TODO: Implement logic
	return false
}
