package linked_lists

import "errors"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type SinglyLinkedList[T comparable] struct {
	Head *Node[T]
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{Head: nil}
}

func (list *SinglyLinkedList[T]) AddItem(val T, pos ...int) error {
	// Determine the position to insert
	position := -1
	if len(pos) > 0 {
		if pos[0] < 1 {
			return errors.New("invalid position")
		}
		position = pos[0]
	}

	node := &Node[T]{
		Value: val,
	}

	// Case: The list is empty
	if list.Head == nil {
		list.Head = node
		return nil
	}

	// Case: Insert at the beginning
	if position == 1 {
		node.Next = list.Head
		list.Head = node
		return nil
	}

	// Case: Insert at the end
	if position == -1 {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
		return nil
	}

	// Case: Insert at the specified position
	current := list.Head
	currentPos := 1

	for current != nil && currentPos < position-1 {
		current = current.Next
		currentPos++
	}

	// If current is nil, position is out of bounds
	if current == nil {
		return errors.New("invalid position")
	}

	node.Next = current.Next
	current.Next = node
	return nil
}

func (list *SinglyLinkedList[T]) GetItems() []T {
	var listItems []T
	item := list.Head

	for item != nil {
		listItems = append(listItems, item.Value)
		item = item.Next
	}

	return listItems
}

func (list *SinglyLinkedList[T]) RemoveItemByValue(val T) error {

	if list.Head == nil {
		return errors.New("empty list")
	}

	// Special case: removing the head node
	if list.Head.Value == val {
		list.Head = list.Head.Next
		return nil
	}

	item := list.Head
	for item.Next != nil {
		if item.Next.Value == val {
			item.Next = item.Next.Next
			return nil
		}
		item = item.Next
	}

	return errors.New("list doesn't contain provided value")
}

func (list *SinglyLinkedList[T]) RemoveItemByPosition(pos int) error {

	if list.Head == nil {
		return errors.New("empty list")
	}

	if pos < 1 {
		return errors.New("invalid position")
	}

	if pos == 1 {
		list.Head = list.Head.Next
		return nil
	}

	item := list.Head
	position := 1

	for item != nil && position < pos-1 {
		item = item.Next
		position++
	}

	if item == nil || item.Next == nil {
		return errors.New("invalid position")
	}

	// Remove the node by skipping it
	item.Next = item.Next.Next

	return nil
}
