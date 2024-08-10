package stacks

import "errors"

type Stack[T any] struct {
	items []T
}

// CreateStack Function to create a stack. It initializes size of stack as 0
func CreateStack[T any]() *Stack[T] {
	return &Stack[T]{items: []T{}}
}

// IsEmpty Stack is empty when stack size is 0
func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

// Push add item to stack
func (stack *Stack[T]) Push(Value T) {
	stack.items = append(stack.items, Value)
}

// Pop remove last item from stack
func (stack *Stack[T]) Pop() (T, error) {
	if stack.IsEmpty() {
		var zero T // zero value for type T
		return zero, errors.New("stack is empty")
	}
	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return lastItem, nil
}

func (stack *Stack[T]) Peek() (any, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return stack.items[len(stack.items)-1], nil
}
