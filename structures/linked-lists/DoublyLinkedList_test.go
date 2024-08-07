package linked_lists

import (
	"testing"
)

func TestAddItem2(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	// Add items at various positions
	if err := list.AddItem(10); err != nil {
		t.Errorf("Error adding item: %v", err)
	}
	if err := list.AddItem(20, 1); err != nil {
		t.Errorf("Error adding item: %v", err)
	}
	if err := list.AddItem(30, 2); err != nil {
		t.Errorf("Error adding item: %v", err)
	}
	if err := list.AddItem(40, 1); err != nil {
		t.Errorf("Error adding item: %v", err)
	}

	expected := []int{40, 20, 30, 10}
	actual := list.GetItems()
	if !equal(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	if err := list.AddItem(666, -10); err == nil {
		t.Errorf("Expected error: pos out of range")
	}

	if err := list.AddItem(661, 0); err == nil {
		t.Errorf("Expected error: pos out of range")
	}

	if err := list.AddItem(626, 55); err == nil {
		t.Errorf("Expected error: pos out of range")
	}
}

func TestRemoveItemByValue2(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.AddItem(10)
	list.AddItem(20)
	list.AddItem(30)
	list.AddItem(40)

	// Remove item by value
	if err := list.RemoveItemByValue(20); err != nil {
		t.Errorf("Error removing item by value: %v", err)
	}

	expected := []int{10, 30, 40}
	actual := list.GetItems()
	if !equal(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Try to remove a non-existing item
	if err := list.RemoveItemByValue(50); err == nil {
		t.Error("Expected an error when removing non-existing item, but got nil")
	}
}

func TestRemoveItemByPosition2(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.AddItem(10)
	list.AddItem(20)
	list.AddItem(30)
	list.AddItem(40)

	// Remove item by position
	if err := list.RemoveItemByPosition(2); err != nil {
		t.Errorf("Error removing item by position: %v", err)
	}

	expected := []int{10, 30, 40}
	actual := list.GetItems()
	if !equal(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Try to remove an item at an invalid position
	if err := list.RemoveItemByPosition(5); err == nil {
		t.Error("Expected an error when removing item at invalid position, but got nil")
	}

	if err := list.RemoveItemByPosition(-69); err == nil {
		t.Error("Expected an error when removing item at invalid position, but got nil")
	}
}

func TestFindItem(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.AddItem(10)
	list.AddItem(20)
	list.AddItem(30)

	// Find an existing item
	node, err := list.FindItem(20)
	if err != nil {
		t.Errorf("Error finding item: %v", err)
	}
	if node == nil || node.Value != 20 {
		t.Errorf("Expected to find node with value 20, but got %v", node)
	}

	// Try to find a non-existing item
	node, err = list.FindItem(50)
	if err == nil {
		t.Error("Expected an error when finding non-existing item, but got nil")
	}
	if node != nil {
		t.Errorf("Expected to not find node, but got %v", node)
	}
}

// Helper function to compare two slices for equality
func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
