package trees

import (
	"bytes"
	"os"
	"testing"
)

// Test for InsertNode, Search, FindMin, FindMax, and DeleteNode
func TestBinarySearchTree(t *testing.T) {
	tree := NewBinarySearchTree()

	tree.InsertNode(&Node{Value: 5})
	tree.InsertNode(&Node{Value: 3})
	tree.InsertNode(&Node{Value: 7})
	tree.InsertNode(&Node{Value: 1})
	tree.InsertNode(&Node{Value: 9})

	if !tree.Search(tree.Root, 5) {
		t.Error("Expected to find node with value 5, but it was not found")
	}
	if !tree.Search(tree.Root, 9) {
		t.Error("Expected to find node with value 9, but it was not found")
	}

	if tree.Search(tree.Root, 100) {
		t.Error("Expected not to find node with value 100, but it was found")
	}

	min := FindMin(tree.Root)
	if min.Value != 1 {
		t.Errorf("Expected min value to be 1, got %d", min.Value)
	}

	max := tree.FindMax()
	if max != 9 {
		t.Errorf("Expected max value to be 9, got %d", max)
	}

	err := tree.DeleteNode(3) // Delete node with one child
	if err != nil {
		t.Error("Error deleting node with value 3:", err)
	}
	if tree.Search(tree.Root, 3) {
		t.Error("Expected node with value 3 to be deleted, but it was found")
	}

	err = tree.DeleteNode(5) // Delete node with two children
	if err != nil {
		t.Error("Error deleting node with value 5:", err)
	}
	if tree.Search(tree.Root, 5) {
		t.Error("Expected node with value 5 to be deleted, but it was found")
	}
}

func captureOutput(f func()) string {
	old := os.Stdout // Keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	os.Stdout = old // Restore original stdout

	return buf.String()
}

// Test for PreOrderTraversal and PostOrderTraversal
func TestTreeTraversals(t *testing.T) {
	tree := NewBinarySearchTree()

	tree.InsertNode(&Node{Value: 5})
	tree.InsertNode(&Node{Value: 3})
	tree.InsertNode(&Node{Value: 7})
	tree.InsertNode(&Node{Value: 1})
	tree.InsertNode(&Node{Value: 9})

	expectedPreOrder := "5\n3\n1\n7\n9\n"
	actualPreOrder := captureOutput(func() {
		tree.PreOrderTraversal(tree.Root)
	})

	if actualPreOrder != expectedPreOrder {
		t.Errorf("Pre-order traversal mismatch. Expected:\n%s\nGot:\n%s", expectedPreOrder, actualPreOrder)
	}

	expectedPostOrder := "1\n3\n9\n7\n5\n"
	actualPostOrder := captureOutput(func() {
		tree.PostOrderTraversal(tree.Root)
	})

	if actualPostOrder != expectedPostOrder {
		t.Errorf("Post-order traversal mismatch. Expected:\n%s\nGot:\n%s", expectedPostOrder, actualPostOrder)
	}
}
