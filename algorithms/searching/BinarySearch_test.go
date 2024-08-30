package searching

import (
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	// Assume array is sorted
	arr := []int{4, 5, 15, 17, 21, 69, 72, 314, 420}

	tests := []struct {
		name     string
		array    []int
		target   int
		expected int
	}{
		{"Found at beginning", arr, 69, 5},
		{"Found in middle", arr, 314, 7},
		{"Found at end", arr, 420, 8},
		{"Not found", arr, 1, -1},
		{"Empty array", []int{}, 69, -1},
		{"Single element - found", []int{69}, 69, 0},
		{"Single element - not found", []int{69}, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Ints(tt.array)
			index := BinarySearch(tt.array, tt.target)
			if index != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, index)
			}
		})
	}
}
