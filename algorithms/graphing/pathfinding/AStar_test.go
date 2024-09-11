package pathfinding

import (
	"dsa-go/structures/graphs"
	"fmt"
	"testing"
)

func TestFindPathWithAStar(t *testing.T) {
	// Create a new weighted graph
	g := graphs.NewWeightedGraph()

	// Adding positive weighted edges
	g.AddEdge(0, graphs.WeightedEdge{Node: 1, Weight: 5}) // A → B
	g.AddEdge(1, graphs.WeightedEdge{Node: 0, Weight: 5}) // B → A (reverse edge)

	g.AddEdge(0, graphs.WeightedEdge{Node: 2, Weight: 10}) // A → C
	g.AddEdge(2, graphs.WeightedEdge{Node: 0, Weight: 10}) // C → A (reverse edge)

	g.AddEdge(1, graphs.WeightedEdge{Node: 3, Weight: 2}) // B → D
	g.AddEdge(3, graphs.WeightedEdge{Node: 1, Weight: 2}) // D → B (reverse edge)

	g.AddEdge(2, graphs.WeightedEdge{Node: 3, Weight: 8}) // C → D
	g.AddEdge(3, graphs.WeightedEdge{Node: 2, Weight: 8}) // D → C (reverse edge)

	g.AddEdge(3, graphs.WeightedEdge{Node: 4, Weight: 4}) // D → E
	g.AddEdge(4, graphs.WeightedEdge{Node: 3, Weight: 4}) // E → D (reverse edge)

	// Additional positive edge
	g.AddEdge(2, graphs.WeightedEdge{Node: 4, Weight: 3}) // C → E
	g.AddEdge(4, graphs.WeightedEdge{Node: 2, Weight: 3}) // E → C (reverse edge)

	// Test case: Find the shortest path from node 0 (A) to node 3 (D)
	expectedPath := []int{0, 1, 3} // A → B → D

	// Run the A* algorithm
	path, err := FindPathWithAStar(g, 0, 3)
	if err != nil {
		t.Fatalf("Error in A* algorithm: %v", err)
	}

	// Check if the path is as expected
	if len(path) != len(expectedPath) {
		t.Errorf("Expected path length %d, but got %d", len(expectedPath), len(path))
		return
	}

	for i, node := range expectedPath {
		if path[i] != node {
			t.Errorf("Expected node %d at position %d, but got %d", node, i, path[i])
		}
	}

	// Print out the graph
	fmt.Println("Graph")
	graphs.PrintWeightedGraph(g)

	// Print the shortest path
	fmt.Println("Shortest path:", path)
}
