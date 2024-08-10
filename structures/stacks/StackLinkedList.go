package stacks

import "errors"

// Implementation of Stack Data Structure using Linked List

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LLStack[T any] struct {
	top *Node[T]
}

// CreateLLStack Function to create a stack. It initializes size of stack as 0
func CreateLLStack[T any]() *LLStack[T] {
	return &LLStack[T]{}
}

// IsEmpty Stack is empty when stack size is 0
func (stack *LLStack[T]) IsEmpty() bool {
	return stack.top == nil
}

// Push add item to stack
func (stack *LLStack[T]) Push(Value T) {
	newNode := &Node[T]{value: Value, next: stack.top}
	stack.top = newNode
}

// Pop remove last item from stack
func (stack *LLStack[T]) Pop() (T, error) {
	if stack.IsEmpty() {
		var zero T // zero value for type T
		return zero, errors.New("stack is empty")
	}
	lastItem := stack.top.value
	stack.top = stack.top.next
	return lastItem, nil
}

func (stack *LLStack[T]) Peek() (any, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return stack.top.value, nil
}
