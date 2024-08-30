package hash_tables

import (
	"testing"
)

// TestNewHashTable verifies that a new hash table is initialized correctly
func TestNewHashTable(t *testing.T) {
	ht := NewHashTable()

	if len(ht.table.entries) != defaultTableSize {
		t.Fatalf("Expected table size %d, got %d", defaultTableSize, len(ht.table.entries))
	}
	if *ht.numEntries != 0 {
		t.Fatalf("Expected numEntries to be 0, got %d", *ht.numEntries)
	}
}

// TestPutAndFind tests the Put and Find methods
func TestPutAndFind(t *testing.T) {
	ht := NewHashTable()

	ht.Put(1, "one")
	ht.Put(2, "two")
	ht.Put(11, "eleven") // This should cause a collision and trigger resizing

	value := ht.Find(1)
	if value != "one" {
		t.Fatalf("Expected 'one', got %s", value)
	}
	value = ht.Find(2)
	if value != "two" {
		t.Fatalf("Expected 'two', got %s", value)
	}
	value = ht.Find(11)
	if value != "eleven" {
		t.Fatalf("Expected 'eleven', got %s", value)
	}
	value = ht.Find(3)
	if value != "Entry not found" {
		t.Fatalf("Expected 'Entry not found', got %s", value)
	}
}

// TestDelete tests the Delete method
func TestDelete(t *testing.T) {
	ht := NewHashTable()

	ht.Put(1, "one")
	ht.Put(2, "two")

	if err := ht.Delete(1); err != nil {
		t.Fatalf("Unexpected error while deleting key 1: %v", err)
	}

	value := ht.Find(1)
	if value != "Entry not found" {
		t.Fatalf("Expected 'Entry not found' after deleting key 1, got %s", value)
	}

	if err := ht.Delete(1); err == nil {
		t.Fatal("Expected error while deleting non-existent key 1, got nil")
	}
}

// TestResize tests if resizing works correctly
func TestResize(t *testing.T) {
	ht := NewHashTable()

	for i := 0; i < 15; i++ {
		ht.Put(i, string(rune('a'+i)))
	}

	if len(ht.table.entries) <= defaultTableSize {
		t.Fatal("Expected table size to increase after adding more entries")
	}

	for i := 0; i < 15; i++ {
		value := ht.Find(i)
		expected := string(rune('a' + i))
		if value != expected {
			t.Fatalf("Expected %s, got %s", expected, value)
		}
	}
}
