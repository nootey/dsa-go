package queue

import (
	"errors"
)

// This is a simple priority queue implementation, without any heap operations

type PriorityItem[T any] struct {
	Item     T
	Priority int
}

type PriorityQueue[T any] struct {
	items []PriorityItem[T]
}

// CreatePriorityQueue initializes a new simple queue with a given capacity
func CreatePriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: []PriorityItem[T]{},
	}
}

// sortByPriority sorts the items based on priority
func (pq *PriorityQueue[T]) sortByPriority() {
	for i := 0; i < len(pq.items)-1; i++ {
		for j := i + 1; j < len(pq.items); j++ {
			if pq.items[i].Priority > pq.items[j].Priority {
				pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
			}
		}
	}
}

// EnqueuePriority adds an item to the priority queue with a given priority
func (pq *PriorityQueue[T]) EnqueuePriority(item T, priority int) error {
	// By default, priority queues using heaps do not guarantee the order of items with the same priority
	if priority < 1 {
		return errors.New("invalid priority")
	}

	pq.items = append(pq.items, PriorityItem[T]{Item: item, Priority: priority})

	pq.sortByPriority()

	return nil
}

// DequeuePriority removes and returns the item with the highest priority
func (pq *PriorityQueue[T]) DequeuePriority() (T, error) {
	if pq.Size() == 0 {
		var zero T
		return zero, errors.New("priority queue is empty")
	}

	// The item with the highest priority is the first in the sorted items list
	item := pq.items[0].Item
	pq.items = pq.items[1:]

	return item, nil
}

// IsEmpty checks if the priority queue is empty
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.items) == 0
}

// Size returns the number of items in the priority queue
func (pq *PriorityQueue[T]) Size() int {
	return len(pq.items)
}
