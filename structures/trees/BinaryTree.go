package trees

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

// InsertNode adds a new node with the given value to the binary tree
func (t *BinaryTree) InsertNode(value int) {
	// TODO: Implement logic
}

// Search returns true if the value is present in the binary tree, otherwise false
func (t *BinaryTree) Search(value int) bool {
	// TODO: Implement logic
	return false
}

// DeleteNode removes a node with the given value from the binary tree
func (t *BinaryTree) DeleteNode(value int) {
	// TODO: Implement logic
}

// PreOrderTraversal prints the nodes of the binary tree in pre-order
func (t *BinaryTree) PreOrderTraversal() {
	// TODO: Implement logic
}

// PostOrderTraversal prints the nodes of the binary tree in post-order
func (t *BinaryTree) PostOrderTraversal() {
	// TODO: Implement logic
}

// FindMin returns the minimum value in the binary tree
func (t *BinaryTree) FindMin() int {
	// TODO: Implement logic
	return 0
}

// FindMax returns the maximum value in the binary tree
func (t *BinaryTree) FindMax() int {
	// TODO: Implement logic
	return 0
}
