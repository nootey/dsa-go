package sorting

import "testing"

func TestHeapSortAscending(t *testing.T) {
	data := []int{4, 10, 3, 5, 1}
	expected := []int{1, 3, 4, 5, 10}

	result := HeapSort(data, false)

	if !equal(result, expected) {
		t.Errorf("HeapSort(data, true) = %v; want %v", result, expected)
	}
}

func TestHeapSortDescending(t *testing.T) {
	data := []int{4, 10, 3, 5, 1}
	expected := []int{10, 5, 4, 3, 1}

	result := HeapSort(data, true)

	if !equal(result, expected) {
		t.Errorf("HeapSort(data, false) = %v; want %v", result, expected)
	}
}

// Helper function to compare slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
