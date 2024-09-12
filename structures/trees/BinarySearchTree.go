package trees

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// InsertNode adds a new node with the given value to the binary tree
func (t *BinarySearchTree) InsertNode(node *Node) {
	t.Root = Insert(t.Root, node)
}

func Insert(root *Node, node *Node) *Node {

	if root == nil {
		root = node
	} else if node.Value < root.Value {
		root.Left = Insert(root.Left, node)
	} else {
		root.Right = Insert(root.Right, node)
	}

	return root

}

func (t *BinarySearchTree) InOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	// left → root → right
	t.InOrderTraversal(root.Left)
	fmt.Println(root.Value)
	t.InOrderTraversal(root.Right)
}

func (t *BinarySearchTree) ReverseInOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	// right → root → left
	t.ReverseInOrderTraversal(root.Right)
	fmt.Println(root.Value)
	t.ReverseInOrderTraversal(root.Left)
}

// PreOrderTraversal prints the nodes of the binary tree in pre-order
func (t *BinarySearchTree) PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	// root → left → right
	fmt.Println(root.Value)
	t.PreOrderTraversal(root.Left)
	t.PreOrderTraversal(root.Right)
}

// PostOrderTraversal prints the nodes of the binary tree in post-order
func (t *BinarySearchTree) PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	// left → right → root
	t.PostOrderTraversal(root.Left)
	t.PostOrderTraversal(root.Right)
	fmt.Println(root.Value)
}

// Search returns true if the value is present in the binary tree, otherwise false
func (t *BinarySearchTree) Search(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if value < root.Value {
		return t.Search(root.Left, value)
	} else if value > root.Value {
		return t.Search(root.Right, value)
	}
	return true
}

// DeleteNode removes a node with the given value from the binary tree
func (t *BinarySearchTree) DeleteNode(value int) error {
	if !t.Search(t.Root, value) {
		return fmt.Errorf("node with value %d does not exist", value)
	}

	t.Root = deleteNode(t.Root, value)
	return nil
}

// Helper function to delete a node from the binary tree
func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	// Traverse the tree to find the node to delete
	if value < root.Value {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteNode(root.Right, value)
	} else {
		// Node found

		// Case 1: No child (leaf node)
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// Case 2: One child
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// Case 3: Two children
		// Find the in-order successor (smallest value in the right subtree)
		minRight := FindMin(root.Right)

		// Replace the current node's value with the successor's value
		root.Value = minRight.Value

		// Delete the in-order successor
		root.Right = deleteNode(root.Right, minRight.Value)
	}

	return root
}

// FindMin returns the minimum value in the binary tree
func FindMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// FindMax returns the maximum value in the binary tree
func (t *BinarySearchTree) FindMax() int {
	if t.Root == nil {
		return -1 // Return some indication if the tree is empty
	}

	// Traverse to the farthest right node
	current := t.Root
	for current.Right != nil {
		current = current.Right
	}
	return current.Value
}
