package arrays

import (
	"testing"
)

func TestNewMDArray(t *testing.T) {
	// Test case 1: Create an MDArray with initial data
	arr := NewMDArray(2, 3, []interface{}{1, 2, 3}, []interface{}{4, 5, 6})
	if len(arr.data) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(arr.data))
	}
	if len(arr.data[0]) != 3 {
		t.Errorf("Expected 3 columns, got %d", len(arr.data[0]))
	}
	if arr.data[0][0] != 1 {
		t.Errorf("Expected 1 at (0,0), got %v", arr.data[0][0])
	}
	if arr.size != 2 {
		t.Errorf("Expected size 2, got %d", arr.size)
	}

	// Test case 2: Create an MDArray without initial data
	arr2 := NewMDArray(3, 4)
	if len(arr2.data) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(arr2.data))
	}
	if len(arr2.data[0]) != 4 {
		t.Errorf("Expected 4 columns, got %d", len(arr2.data[0]))
	}
}

func TestGetItem(t *testing.T) {
	arr := NewMDArray(2, 3, []interface{}{1, 2, 3}, []interface{}{4, 5, 6})

	// Test valid index
	row, err := arr.GetItem(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if row[0] != 4 {
		t.Errorf("Expected 4, got %v", row[0])
	}

	// Test invalid index
	_, err = arr.GetItem(3)
	if err == nil {
		t.Errorf("Expected error for out-of-bounds index")
	}

	_, err = arr.GetItem(-69)
	if err == nil {
		t.Errorf("Expected error for out-of-bounds index")
	}
}

func TestAddItem(t *testing.T) {
	arr := NewMDArray(2, 3)

	// Test appending an item
	err := arr.AddItem([]interface{}{1, 2, 3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(arr.data) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(arr.data))
	}

	// Test inserting an item at a specific index
	err = arr.AddItem([]interface{}{4, 5, 6}, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(arr.data) != 4 {
		t.Errorf("Expected 4 rows, got %d", len(arr.data))
	}
	if arr.data[0][0] != 4 {
		t.Errorf("Expected 4 at (0,0), got %v", arr.data[0][0])
	}

	// Test out-of-bounds index
	err = arr.AddItem([]interface{}{7, 8, 9}, 3)
	if err == nil {
		t.Errorf("Expected error for out-of-bounds index")
	}

}

func TestRemoveItem(t *testing.T) {
	arr := NewMDArray(3, 3, []interface{}{1, 2, 3}, []interface{}{4, 5, 6}, []interface{}{7, 8, 9})

	// Test removing an item
	err := arr.RemoveItem(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(arr.data) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(arr.data))
	}
	if arr.data[1][0] != 7 {
		t.Errorf("Expected 7 at (1,0), got %v", arr.data[1][0])
	}

	// Test removing an out-of-bounds index
	err = arr.RemoveItem(2)
	if err == nil {
		t.Errorf("Expected error for out-of-bounds index")
	}
}
