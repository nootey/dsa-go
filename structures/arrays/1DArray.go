package arrays

import (
	"errors"
)

type SDArray struct {
	data     []int
	size     int
	capacity int
}

func (a *SDArray) Size() int {
	return a.size
}

func (a *SDArray) Capacity() int {
	return a.capacity
}

func NewSDArray(initialCapacity int) *SDArray {
	return &SDArray{
		data:     make([]int, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

func (a *SDArray) Add(element int) {
	if a.size == a.capacity {
		a.resize()
	}
	a.data[a.size] = element
	a.size++
}

// Doubling the capacity is a common strategy to minimize the frequency of resizing operations.
func (a *SDArray) resize() {
	newCapacity := a.capacity * 2
	newData := make([]int, newCapacity)
	copy(newData, a.data)
	a.data = newData
	a.capacity = newCapacity
}

func (a *SDArray) Get(index int) (int, error) {
	if index < 0 || index >= a.size {
		return 0, errors.New("index out of bounds")
	}
	return a.data[index], nil
}

func (a *SDArray) Set(index int, element int) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of bounds")
	}
	a.data[index] = element
	return nil
}

func (a *SDArray) Remove(index int) error {
	if index < 0 || index >= a.size {
		return errors.New("index out of bounds")
	}
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	return nil
}
