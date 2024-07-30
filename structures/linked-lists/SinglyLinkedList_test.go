package linked_lists

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	// Test adding to an empty list
	if err := list.AddItem(1); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if items := list.GetItems(); len(items) != 1 || items[0] != 1 {
		t.Errorf("Expected [1], got %v", items)
	}

	// Test adding at the beginning
	if err := list.AddItem(2, 1); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if items := list.GetItems(); len(items) != 2 || items[0] != 2 || items[1] != 1 {
		t.Errorf("Expected [2, 1], got %v", items)
	}

	// Test adding in the middle
	if err := list.AddItem(3, 2); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if items := list.GetItems(); len(items) != 3 || items[0] != 2 || items[1] != 3 || items[2] != 1 {
		t.Errorf("Expected [2, 3, 1], got %v", items)
	}

	// Test adding at an invalid position
	if err := list.AddItem(4, 10); err == nil {
		t.Errorf("Expected error for invalid position, got nil")
	}
}

func TestRemoveItemByValue(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.AddItem(1)
	list.AddItem(2)
	list.AddItem(3)

	// Test removing an item
	if err := list.RemoveItemByValue(2); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if items := list.GetItems(); len(items) != 2 || items[0] != 1 || items[1] != 3 {
		t.Errorf("Expected [1, 3], got %v", items)
	}

	// Test removing an item that doesn't exist
	if err := list.RemoveItemByValue(4); err == nil {
		t.Errorf("Expected error for removing non-existent item, got nil")
	}
}

func TestRemoveItemByPosition(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.AddItem(1)
	list.AddItem(2)
	list.AddItem(3)

	// Test removing an item by position
	if err := list.RemoveItemByPosition(2); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if items := list.GetItems(); len(items) != 2 || items[0] != 1 || items[1] != 3 {
		t.Errorf("Expected [1, 3], got %v", items)
	}

	// Test removing an item at invalid position
	if err := list.RemoveItemByPosition(5); err == nil {
		t.Errorf("Expected error for invalid position, got nil")
	}
}

func TestRemoveItemByPosition_EmptyList(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	// Test removing an item by position from an empty list
	if err := list.RemoveItemByPosition(1); err == nil {
		t.Errorf("Expected error for empty list, got nil")
	}
}

func TestAddItem_AtInvalidPosition(t *testing.T) {
	list := NewSinglyLinkedList[int]()

	// Test adding item at invalid position
	if err := list.AddItem(1, -1); err == nil {
		t.Errorf("Expected error for invalid position, got nil")
	}
	if err := list.AddItem(1, 0); err == nil {
		t.Errorf("Expected error for invalid position, got nil")
	}
}
