package sorting

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	array := []int{69, 420, 21, 15, 4, 17, 314, 5, 72}
	SelectionSort(array)

	if !reflect.DeepEqual([]int{4, 5, 15, 17, 21, 69, 72, 314, 420}, array) {
		t.Errorf("Expected [4, 5, 15, 17, 21, 69, 72, 314, 420], got %v", array)
	}
}
