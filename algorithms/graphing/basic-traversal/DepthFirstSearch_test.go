package basic_traversal

import (
	"bytes"
	"dsa-go/structures/graphs"
	"fmt"
	"log"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	// Create a pipe to capture the output
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	// Run the function
	f()

	// Close the writer to flush the buffer
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestDepthFirstSearchRecursive(t *testing.T) {
	g, err := graphs.CreateNewGraph(5)
	if err != nil {
		log.Fatal(err)
	}
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddNode("E")

	g.AddEdge(0, 1) // A → B
	g.AddEdge(1, 2) // B → C
	g.AddEdge(1, 4) // B → E
	g.AddEdge(2, 3) // C → D
	g.AddEdge(2, 4) // C → E
	g.AddEdge(4, 0) // E → A
	g.AddEdge(4, 2) // E → C

	expectedOutput := "Visited: A\nVisited: B\nVisited: C\nVisited: D\nVisited: E\n"
	actualOutput := captureOutput(func() {
		err := DepthFirstSearchRecursive(g, 0)
		if err != nil {
			log.Fatal(err)
		}
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}

	fmt.Println(actualOutput)
}

func TestDepthFirstSearchStack(t *testing.T) {
	g, err := graphs.CreateNewGraph(5)
	if err != nil {
		log.Fatal(err)
	}
	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddNode("D")
	g.AddNode("E")

	g.AddEdge(0, 1) // A → B
	g.AddEdge(1, 2) // B → C
	g.AddEdge(1, 4) // B → E
	g.AddEdge(2, 3) // C → D
	g.AddEdge(2, 4) // C → E
	g.AddEdge(4, 0) // E → A
	g.AddEdge(4, 2) // E → C

	expectedOutput := "Visited: A\nVisited: B\nVisited: E\nVisited: C\nVisited: D\n"
	actualOutput := captureOutput(func() {
		err := DepthFirstSearchStack(g, 0)
		if err != nil {
			t.Errorf("Error during DFS stack traversal: %v", err)
		}
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}

	fmt.Println(actualOutput)
}

func TestDepthFirstSearchStackWithError(t *testing.T) {

	// Not adding nodes or edges
	g, err := graphs.CreateNewGraph(5)
	if err != nil {
		log.Fatal(err)
	}

	err = DepthFirstSearchRecursive(g, 0)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}

	err = DepthFirstSearchStack(g, 0)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}

	// Provide invalid source index
	err = DepthFirstSearchRecursive(g, -1)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}

	err = DepthFirstSearchStack(g, -1)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}

}
