package queue

import (
	"errors"
)

type SimpleQueue[T any] struct {
	front    int
	rear     int
	capacity int
	items    []T
}

// CreateSimpleQueue initializes a new simple queue with a given capacity
func CreateSimpleQueue[T any](capacity int) *SimpleQueue[T] {
	if capacity < 1 {
		return nil
	}

	return &SimpleQueue[T]{
		front:    0,
		rear:     -1,
		capacity: capacity,
		items:    make([]T, 0, capacity),
	}
}

func (q *SimpleQueue[T]) EnqueueSimple(item T) error {
	if q.rear == q.capacity-1 {
		return errors.New("queue is full")
	}

	q.rear++
	q.items = append(q.items, item)

	return nil
}

func (q *SimpleQueue[T]) DequeueSimple() (T, error) {
	var zeroValue T // default value for type T

	if q.IsEmpty() {
		return zeroValue, errors.New("queue is empty")
	}

	item := q.items[0]
	q.items = q.items[1:]
	q.rear--

	return item, nil
}

func (q *SimpleQueue[T]) IsEmpty() bool {
	return q.front > q.rear
}

func (q *SimpleQueue[T]) Size() int {
	return len(q.items)
}
