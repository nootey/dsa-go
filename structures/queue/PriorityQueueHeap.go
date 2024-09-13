package queue

import (
	"dsa-go/structures/heap"
	"errors"
)

type PriorityHeapItem[T any] struct {
	Item     T
	Priority int
}

type PriorityHeapQueue[T any] struct {
	heap        *heap.Heap
	priorityMap map[int][]T
}

// CreatePriorityHeapQueue initializes a new priority queue that uses a heap
func CreatePriorityHeapQueue[T any](isMinHeap bool) *PriorityHeapQueue[T] {
	return &PriorityHeapQueue[T]{
		heap:        &heap.Heap{IsMinHeap: isMinHeap, Data: []int{}},
		priorityMap: make(map[int][]T),
	}
}

// EnqueuePriority adds an item to the priority queue with a given priority
func (pq *PriorityHeapQueue[T]) EnqueuePriority(item T, priority int) {
	pq.heap.Insert(priority)
	pq.priorityMap[priority] = append(pq.priorityMap[priority], item)
}

// DequeuePriority removes and returns the item with the highest or lowest priority
func (pq *PriorityHeapQueue[T]) DequeuePriority() (T, error) {
	if pq.Size() == 0 {
		var zero T
		return zero, errors.New("priority queue is empty")
	}

	rootPriority := pq.heap.ExtractRoot()
	items := pq.priorityMap[rootPriority]

	// Remove the first item (FIFO)
	dequeuedItem := items[0]

	if len(items) > 1 {
		pq.priorityMap[rootPriority] = items[1:]
	} else {
		delete(pq.priorityMap, rootPriority) // No more items for this priority
	}

	return dequeuedItem, nil
}

// IsEmpty checks if the priority queue is empty
func (pq *PriorityHeapQueue[T]) IsEmpty() bool {
	return pq.Size() == 0
}

// Size returns the number of items in the priority queue
func (pq *PriorityHeapQueue[T]) Size() int {
	return len(pq.heap.Data)
}
