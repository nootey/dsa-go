package queue

import "testing"

func TestPriorityHeapQueue(t *testing.T) {
	// Create a new priority queue (min-heap)
	pq := CreatePriorityHeapQueue[string](true)

	// Test EnqueuePriority
	pq.EnqueuePriority("low-priority", 10)
	pq.EnqueuePriority("medium-priority", 5)
	pq.EnqueuePriority("high-priority", 1)

	// Test DequeuePriority
	item, err := pq.DequeuePriority()
	if err != nil {
		t.Fatalf("DequeuePriority() returned an error: %v", err)
	}
	if item != "high-priority" {
		t.Errorf("Expected 'high-priority', got '%s'", item)
	}

	// Test DequeuePriority with remaining items
	item, err = pq.DequeuePriority()
	if err != nil {
		t.Fatalf("DequeuePriority() returned an error: %v", err)
	}
	if item != "medium-priority" {
		t.Errorf("Expected 'medium-priority', got '%s'", item)
	}

	item, err = pq.DequeuePriority()
	if err != nil {
		t.Fatalf("DequeuePriority() returned an error: %v", err)
	}
	if item != "low-priority" {
		t.Errorf("Expected 'low-priority', got '%s'", item)
	}

	// Test DequeuePriority on empty queue
	item, err = pq.DequeuePriority()
	if err == nil {
		t.Fatal("Expected an error when dequeuing from an empty queue, but got none")
	}
	if item != "" {
		t.Errorf("Expected an empty string, got '%s'", item)
	}

	// Test IsEmpty
	if !pq.IsEmpty() {
		t.Error("Expected queue to be empty after all items are dequeued")
	}

	// Test Size
	if size := pq.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Add more items
	pq.EnqueuePriority("new-item-1", 7)
	pq.EnqueuePriority("new-item-2", 2)

	// Test Size after adding items
	if size := pq.Size(); size != 2 {
		t.Errorf("Expected size 2, got %d", size)
	}

	// Test DequeuePriority with new items
	item, err = pq.DequeuePriority()
	if err != nil {
		t.Fatalf("DequeuePriority() returned an error: %v", err)
	}
	if item != "new-item-2" {
		t.Errorf("Expected 'new-item-2', got '%s'", item)
	}

	item, err = pq.DequeuePriority()
	if err != nil {
		t.Fatalf("DequeuePriority() returned an error: %v", err)
	}
	if item != "new-item-1" {
		t.Errorf("Expected 'new-item-1', got '%s'", item)
	}

	// Test IsEmpty after all items are dequeued
	if !pq.IsEmpty() {
		t.Error("Expected queue to be empty after all new items are dequeued")
	}
}
