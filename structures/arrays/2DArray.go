package arrays

import (
	"errors"
)

type MDArray struct {
	numRows  int
	numCols  int
	size     int
	capacity int
	data     [][]interface{}
}

func NewMDArray(numRows int, numCols int, initialData ...[]interface{}) *MDArray {

	arr := &MDArray{
		numRows:  numRows,
		numCols:  numCols,
		size:     0,
		capacity: numRows,
		data:     make([][]interface{}, numRows),
	}

	// Initialize each row in the slice
	for i := range arr.data {
		arr.data[i] = make([]interface{}, numCols)
	}

	// Fill the array with data, if provided
	if len(initialData) > 0 {
		for row := 0; row < numRows && row < len(initialData); row++ {
			for col := 0; col < numCols && col < len(initialData[row]); col++ {
				arr.data[row][col] = initialData[row][col]
			}
		}
		arr.size = len(initialData)
	}

	return arr
}

func (arr *MDArray) resize() {
	newCapacity := arr.capacity * 2
	newData := make([][]interface{}, len(arr.data), newCapacity)
	copy(newData, arr.data)
	arr.data = newData
	arr.capacity = newCapacity
}

func (arr *MDArray) GetItem(rowIndex int) ([]interface{}, error) {
	if rowIndex >= arr.size || rowIndex < 0 {
		return nil, errors.New("index out of bounds")
	}

	return arr.data[rowIndex], nil
}

func (arr *MDArray) AddItem(item []interface{}, rowIndex ...int) error {

	// Determine the row index
	index := len(arr.data) // Default to adding at the end
	if len(rowIndex) > 0 {
		index = rowIndex[0]
	}

	// Validate index
	if index < 0 || (index >= arr.size && arr.size > 0) {
		return errors.New("index out of bounds")
	}

	// Ensure capacity
	if arr.size == arr.capacity {
		arr.resize()
	}

	// Shift elements and insert
	if index == len(arr.data) {
		// Appending to the end
		arr.data = append(arr.data, item)
	} else {
		// Insert at specified index
		arr.data = append(arr.data[:index+1], arr.data[index:]...)
		arr.data[index] = item
	}

	arr.numRows++
	arr.size++

	return nil

}

func (arr *MDArray) RemoveItem(index int) error {

	if index >= arr.size || index < 0 {
		return errors.New("index out of bounds")
	}

	if index == len(arr.data)-1 {
		arr.data = arr.data[:len(arr.data)-1]
	} else {
		arr.data = append(arr.data[:index], arr.data[index+1:]...)
	}

	arr.size--
	arr.numRows--

	return nil
}
