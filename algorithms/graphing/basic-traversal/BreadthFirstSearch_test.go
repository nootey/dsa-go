package basic_traversal

import (
	"dsa-go/structures/graphs"
	"fmt"
	"log"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
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

	expectedOutput := "Visited: A\nVisited: B\nVisited: C\nVisited: E\nVisited: D\n"
	actualOutput := captureOutput(func() {
		err := BreadthFirstSearch(g, 0)
		if err != nil {
			log.Fatal(err)
		}
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}

	fmt.Println(actualOutput)

}

func TestBreadthFirstSearchWithError(t *testing.T) {
	// Not adding nodes or edges
	g, err := graphs.CreateNewGraph(5)
	if err != nil {
		log.Fatal(err)
	}

	err = BreadthFirstSearch(g, 0)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}

	// Provide invalid source index
	err = BreadthFirstSearch(g, -1)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}
