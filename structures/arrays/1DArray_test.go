package arrays

import (
	"testing"
)

func TestMyArray_Add(t *testing.T) {
	arr := NewSDArray(2)
	arr.Add(1)
	arr.Add(2)
	if arr.Size() != 2 {
		t.Errorf("Expected size 2, got %d", arr.Size())
	}
	if arr.Capacity() != 2 {
		t.Errorf("Expected capacity 2, got %d", arr.Capacity())
	}
	arr.Add(3)
	if arr.Capacity() != 4 {
		t.Errorf("Expected capacity 4, got %d", arr.Capacity())
	}
}

func TestMyArray_GetSet(t *testing.T) {
	arr := NewSDArray(2)
	arr.Add(1)
	arr.Add(2)
	if val, _ := arr.Get(1); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	arr.Set(1, 3)
	if val, _ := arr.Get(1); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

func TestMyArray_Remove(t *testing.T) {
	arr := NewSDArray(3)
	arr.Add(1)
	arr.Add(2)
	arr.Add(3)
	arr.Remove(1)
	if val, _ := arr.Get(1); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	if arr.Size() != 2 {
		t.Errorf("Expected size 2, got %d", arr.Size())
	}
}
