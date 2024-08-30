package searching

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	// Table-driven tests
	tests := []struct {
		name     string
		array    []int
		target   int
		expected int
	}{
		{"Found at beginning", []int{69, 420, 21, 15, 4, 17, 314, 5, 72}, 69, 0},
		{"Found in middle", []int{69, 420, 21, 15, 4, 17, 314, 5, 72}, 314, 6},
		{"Found at end", []int{69, 420, 21, 15, 4, 17, 314, 5, 72}, 72, 8},
		{"Not found", []int{69, 420, 21, 15, 4, 17, 314, 5, 72}, 1, -1},
		{"Empty array", []int{}, 69, -1},
		{"Single element - found", []int{69}, 69, 0},
		{"Single element - not found", []int{69}, 1, -1},
		{"Multiple occurrences", []int{69, 420, 21, 69, 4, 69, 314, 5, 72}, 69, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := LinearSearch(tt.array, tt.target)
			if index != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, index)
			}
		})
	}
}
