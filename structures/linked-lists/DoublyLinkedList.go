package linked_lists

import "errors"

type DoublyLinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{Head: nil}
}

func (list *DoublyLinkedList[T]) GetItems() []T {
	var listItems []T
	item := list.Head

	for item != nil {
		listItems = append(listItems, item.Value)
		item = item.Next
	}

	return listItems
}

func (list *DoublyLinkedList[T]) AddItem(val T, pos ...int) error {
	// Determine the position to insert
	position := -1
	if len(pos) > 0 {
		if pos[0] < 1 {
			return errors.New("invalid position")
		}
		position = pos[0]
	}

	newNode := &Node[T]{
		Value: val,
	}

	// Case: The list is empty
	if list.Head == nil {
		list.Head = newNode
		return nil
	}

	// Case: Insert at the beginning
	if position == 1 {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
		return nil
	}

	// Case: Insert at the specified position or at the end
	current := list.Head
	currentPos := 1

	// Traverse the list and find the correct positional node
	for current.Next != nil && (position == -1 || currentPos < position-1) {
		current = current.Next
		currentPos++
	}

	if position != -1 && currentPos < position-1 {
		return errors.New("invalid position")
	}

	newNode.Next = current.Next
	newNode.Prev = current
	if current.Next != nil {
		current.Next.Prev = newNode
	} else {
		// Update the tail if the node is inserted at the end
		list.Tail = newNode
	}
	current.Next = newNode

	return nil
}

func (list *DoublyLinkedList[T]) RemoveItemByValue(val T) error {

	if list.Head == nil {
		return errors.New("empty list")
	}

	// Case: remove the first item
	if list.Head.Value == val {
		if list.Head.Next != nil {
			list.Head = list.Head.Next
			list.Head.Prev = nil
		} else {
			// If there's only one node
			list.Head = nil
			list.Tail = nil
		}
		return nil
	}

	item := list.Head
	for item != nil {
		if item.Value == val {

			if item.Next != nil {
				item.Next.Prev = item.Prev
			} else {
				// Update tail if last node is removed
				list.Tail = item.Prev
			}

			if item.Prev != nil {
				item.Prev.Next = item.Next
			}

			return nil
		}
		item = item.Next
	}

	return errors.New("list doesn't contain provided value")
}

func (list *DoublyLinkedList[T]) RemoveItemByPosition(pos int) error {

	if list.Head == nil {
		return errors.New("empty list")
	}

	if pos < 1 {
		return errors.New("invalid position")
	}

	// Case: remove the first item
	if pos == 1 {
		if list.Head.Next != nil {
			list.Head = list.Head.Next
			list.Head.Prev = nil
		} else {
			// If there's only one node
			list.Head = nil
			list.Tail = nil
		}
		return nil
	}

	item := list.Head
	position := 1

	// Find the item by position
	for item != nil && position < pos-1 {
		item = item.Next
		position++
	}

	if item == nil || item.Next == nil {
		return errors.New("invalid position")
	}

	// Remove the node
	nodeToRemove := item.Next
	item.Next = nodeToRemove.Next

	if nodeToRemove.Next != nil {
		nodeToRemove.Next.Prev = item
	} else {
		// Update the tail
		list.Tail = item
	}

	return nil
}

func (list *DoublyLinkedList[T]) FindItem(val T) (*Node[T], error) {

	if list.Head == nil {
		return nil, errors.New("empty list")
	}

	item := list.Head
	for item != nil {
		if item.Value == val {
			return item, nil
		}
		item = item.Next
	}

	return nil, errors.New("list doesn't contain provided value")
}
