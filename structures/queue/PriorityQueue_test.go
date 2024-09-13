package queue

import (
	"testing"
)

func TestPriorityOrdering(t *testing.T) {
	priorityQueue := CreatePriorityQueue[string]()

	priorityQueue.EnqueuePriority("low", 3)
	priorityQueue.EnqueuePriority("medium", 2)
	priorityQueue.EnqueuePriority("high", 1)

	// Expected order: high, medium, low
	expectedOrder := []string{"high", "medium", "low"}
	for _, expected := range expectedOrder {
		item, err := priorityQueue.DequeuePriority()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if item != expected {
			t.Errorf("Expected %v, but got %v", expected, item)
		}
	}
}

func TestPriorityWithSameValues(t *testing.T) {
	priorityQueue := CreatePriorityQueue[string]()

	priorityQueue.EnqueuePriority("first", 1)
	priorityQueue.EnqueuePriority("second", 1)

	// Both items have the same priority
	// They should be dequeued in the order they were added
	expectedOrder := []string{"first", "second"}
	for _, expected := range expectedOrder {
		item, err := priorityQueue.DequeuePriority()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if item != expected {
			t.Errorf("Expected %v, but got %v", expected, item)
		}
	}
}

func TestEnqueueWithZeroPriority(t *testing.T) {
	priorityQueue := CreatePriorityQueue[string]()

	if err := priorityQueue.EnqueuePriority("item", -1); err == nil {
		t.Errorf("Expected error for zero priority, but got none")
	}
}

func TestEnqueueAfterDequeue(t *testing.T) {
	priorityQueue := CreatePriorityQueue[int]()

	priorityQueue.EnqueuePriority(1, 1)
	priorityQueue.EnqueuePriority(2, 3)
	priorityQueue.DequeuePriority() // remove item with priority 1

	// Enqueue again after dequeue
	priorityQueue.EnqueuePriority(3, 2)

	if pqSize := priorityQueue.Size(); pqSize != 2 {
		t.Errorf("Expected size 2, but got %d", pqSize)
	}

	// Expected order: 3, 2
	expectedOrder := []int{3, 2}
	for _, expected := range expectedOrder {
		item, err := priorityQueue.DequeuePriority()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if item != expected {
			t.Errorf("Expected %v, but got %v", expected, item)
		}
	}
}

func TestSizeMethod(t *testing.T) {
	priorityQueue := CreatePriorityQueue[int]()

	priorityQueue.EnqueuePriority(1, 1)
	if pqSize := priorityQueue.Size(); pqSize != 1 {
		t.Errorf("Expected size 1, but got %d", pqSize)
	}

	priorityQueue.EnqueuePriority(2, 2)
	if pqSize := priorityQueue.Size(); pqSize != 2 {
		t.Errorf("Expected size 2, but got %d", pqSize)
	}

	priorityQueue.DequeuePriority()
	if pqSize := priorityQueue.Size(); pqSize != 1 {
		t.Errorf("Expected size 1, but got %d", pqSize)
	}

	priorityQueue.DequeuePriority()
	if pqSize := priorityQueue.Size(); pqSize != 0 {
		t.Errorf("Expected size 0, but got %d", pqSize)
	}
}

func TestIsEmptyMethod(t *testing.T) {
	priorityQueue := CreatePriorityQueue[string]()
	if !priorityQueue.IsEmpty() {
		t.Errorf("Expected priority queue to be empty")
	}

	priorityQueue.EnqueuePriority("item", 1)
	if priorityQueue.IsEmpty() {
		t.Errorf("Expected priority queue to not be empty")
	}

	priorityQueue.DequeuePriority()
	if !priorityQueue.IsEmpty() {
		t.Errorf("Expected priority queue to be empty")
	}

	_, err := priorityQueue.DequeuePriority()
	if err == nil {
		t.Errorf("Expected priority queue to be empty")
	}
}
