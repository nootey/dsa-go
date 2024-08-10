package stacks

import (
	"testing"
)

func TestCreateLLStack(t *testing.T) {
	stack := CreateLLStack[int]()
	if stack.IsEmpty() != true {
		t.Errorf("Expected stack to be empty")
	}
}

func TestLLPush(t *testing.T) {
	stack := CreateLLStack[int]()
	stack.Push(1)
	if stack.IsEmpty() != false {
		t.Errorf("Expected stack to be non-empty after push")
	}
	if top, _ := stack.Peek(); top != 1 {
		t.Errorf("Expected top item to be 1, got %v", top)
	}
}

func TestLLPop(t *testing.T) {
	stack := CreateLLStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item, err := stack.Pop()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if item != 3 {
		t.Errorf("Expected popped item to be 3, got %v", item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if item != 2 {
		t.Errorf("Expected popped item to be 2, got %v", item)
	}

	item, err = stack.Pop()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if item != 1 {
		t.Errorf("Expected popped item to be 1, got %v", item)
	}

	_, err = stack.Pop()
	if err == nil {
		t.Errorf("Expected error when popping from empty stack")
	}
}

func TestLLPeek(t *testing.T) {
	stack := CreateLLStack[int]()
	stack.Push(1)
	stack.Push(2)

	item, err := stack.Peek()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if item != 2 {
		t.Errorf("Expected top item to be 2, got %v", item)
	}

	_, err = stack.Pop()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	item, err = stack.Peek()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if item != 1 {
		t.Errorf("Expected top item to be 1, got %v", item)
	}

	_, err = stack.Pop()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, err = stack.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking from empty stack")
	}
}
