package heap

import (
	"testing"
)

// TestMinHeapInsertAndExtract tests insertion and extraction in a min-heap
func TestMinHeapInsertAndExtract(t *testing.T) {
	h := &Heap{IsMinHeap: true}

	h.Insert(10)
	h.Insert(4)
	h.Insert(15)
	h.Insert(1)

	min := h.GetMin()
	if min != 1 {
		t.Errorf("Expected minimum value to be 1, but got %d", min)
	}

	extractedMin := h.ExtractRoot()
	if extractedMin != 1 {
		t.Errorf("Expected to extract 1, but got %d", extractedMin)
	}

	min = h.GetMin()
	if min != 4 {
		t.Errorf("Expected minimum value to be 4 after extracting 1, but got %d", min)
	}

	expectedValues := []int{4, 10, 15}
	for _, expected := range expectedValues {
		extracted := h.ExtractRoot()
		if extracted != expected {
			t.Errorf("Expected %d, but got %d", expected, extracted)
		}
	}
}

// TestMaxHeapInsertAndExtract tests insertion and extraction in a max-heap
func TestMaxHeapInsertAndExtract(t *testing.T) {
	h := &Heap{IsMinHeap: false}

	h.Insert(10)
	h.Insert(4)
	h.Insert(15)
	h.Insert(1)

	max := h.GetMax()
	if max != 15 {
		t.Errorf("Expected maximum value to be 15, but got %d", max)
	}

	extractedMax := h.ExtractRoot()
	if extractedMax != 15 {
		t.Errorf("Expected to extract 15, but got %d", extractedMax)
	}

	max = h.GetMax()
	if max != 10 {
		t.Errorf("Expected maximum value to be 10 after extracting 15, but got %d", max)
	}

	expectedValues := []int{10, 4, 1}
	for _, expected := range expectedValues {
		extracted := h.ExtractRoot()
		if extracted != expected {
			t.Errorf("Expected %d, but got %d", expected, extracted)
		}
	}
}

// TestBuildHeap tests building a heap from an unsorted slice
func TestBuildHeap(t *testing.T) {
	data := []int{10, 4, 15, 1}

	minHeap := &Heap{IsMinHeap: true}
	minHeap.BuildHeap(data)

	min := minHeap.GetMin()
	if min != 1 {
		t.Errorf("Expected minimum value to be 1 in the built min-heap, but got %d", min)
	}

	maxHeap := &Heap{IsMinHeap: false}
	maxHeap.BuildHeap(data)

	max := maxHeap.GetMax()
	if max != 15 {
		t.Errorf("Expected maximum value to be 15 in the built max-heap, but got %d", max)
	}
}

// TestEmptyHeap tests behavior on an empty heap
func TestEmptyHeap(t *testing.T) {
	h := &Heap{IsMinHeap: true}

	extracted := h.ExtractRoot()
	if extracted != 0 {
		t.Errorf("Expected to extract 0 from an empty heap, but got %d", extracted)
	}

	min := h.GetMin()
	if min != 0 {
		t.Errorf("Expected minimum value to be 0 in an empty heap, but got %d", min)
	}

	max := h.GetMax()
	if max != 0 {
		t.Errorf("Expected maximum value to be 0 in an empty heap, but got %d", max)
	}
}
