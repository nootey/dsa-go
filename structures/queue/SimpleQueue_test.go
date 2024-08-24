package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {

	simpleQueue := CreateSimpleQueue[string](3)

	// Test adding to an empty list
	if err := simpleQueue.EnqueueSimple("abc"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test exceeding capacity
	if err := simpleQueue.EnqueueSimple("def"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := simpleQueue.EnqueueSimple("ghj"); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := simpleQueue.EnqueueSimple("ikl"); err == nil {
		t.Errorf("Unexpected error: %v", err)
	}

}

func TestDequeue(t *testing.T) {

	simpleQueue := CreateSimpleQueue[int](3)
	simpleQueue.EnqueueSimple(1)
	simpleQueue.EnqueueSimple(2)
	simpleQueue.EnqueueSimple(3)

	// Test dequeue
	for simpleQueue.Size() > 0 {
		if err := simpleQueue.DequeueSimple(); err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	// Test dequeue with empty queue
	if err := simpleQueue.DequeueSimple(); err == nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
