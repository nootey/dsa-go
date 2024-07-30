package linked_lists

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

// AddItem Insert value to linked list based on pos (true = beginning, false = end)
func (list *SinglyLinkedList[T]) AddItem(val T, pos bool) {
	node := &Node[T]{
		Value: val,
	}

	if list.Head == nil {
		list.Head = node
		return
	}

	if pos {
		node.Next = list.Head
		list.Head = node
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
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

func (list *SinglyLinkedList[T]) RemoveItem(val T, pos bool) {
	node := &Node[T]{
		Value: val,
	}

	if list.Head == nil {
		list.Head = node
		return
	}

	if pos {
		node.Next = list.Head
		list.Head = node
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}
