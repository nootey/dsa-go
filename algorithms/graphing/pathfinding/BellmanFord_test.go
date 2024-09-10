package pathfinding

import (
	"dsa-go/structures/graphs"
	"fmt"
	"testing"
)

func createTestGraph() *graphs.WeightedGraph {
	g := graphs.NewWeightedGraph()

	// Adding positive weighted edges
	g.AddEdge(0, graphs.WeightedEdge{Node: 1, Weight: 5})
	g.AddEdge(1, graphs.WeightedEdge{Node: 0, Weight: 5})

	g.AddEdge(0, graphs.WeightedEdge{Node: 2, Weight: 10})
	g.AddEdge(2, graphs.WeightedEdge{Node: 0, Weight: 10})

	g.AddEdge(1, graphs.WeightedEdge{Node: 3, Weight: 2})
	g.AddEdge(3, graphs.WeightedEdge{Node: 1, Weight: 2})

	g.AddEdge(2, graphs.WeightedEdge{Node: 3, Weight: 8})
	g.AddEdge(3, graphs.WeightedEdge{Node: 2, Weight: 8})

	g.AddEdge(3, graphs.WeightedEdge{Node: 4, Weight: 4})
	g.AddEdge(4, graphs.WeightedEdge{Node: 3, Weight: 4})

	g.AddEdge(1, graphs.WeightedEdge{Node: 4, Weight: -3})

	g.AddEdge(2, graphs.WeightedEdge{Node: 4, Weight: 3})
	g.AddEdge(4, graphs.WeightedEdge{Node: 2, Weight: 3})

	return g
}

func TestBellmanFord(t *testing.T) {
	g := createTestGraph()
	distances, err := FindPathWithBellmanFord(g, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedDistances := map[int]float64{
		0: 0,
		1: 5,
		2: 5,
		3: 6,
		4: 2,
	}

	for node, expected := range expectedDistances {
		if distances[node] != expected {
			t.Errorf("expected distance to node %d: %f, got: %f", node, expected, distances[node])
		}
	}

	fmt.Println("Graph")
	graphs.PrintWeightedGraph(g)
	fmt.Println("Distance table")
	fmt.Println(distances)
	fmt.Printf("\n")
}

func TestBellmanFordNegativeCycle(t *testing.T) {
	// Create a graph with a negative weight cycle
	g := graphs.NewWeightedGraph()
	g.AddEdge(0, graphs.WeightedEdge{Node: 1, Weight: 1})
	g.AddEdge(1, graphs.WeightedEdge{Node: 2, Weight: -1})
	g.AddEdge(2, graphs.WeightedEdge{Node: 0, Weight: -1})

	_, err := FindPathWithBellmanFord(g, 0)
	if err == nil {
		t.Fatal("expected an error due to negative weight cycle, got nil")
	} else if err.Error() != "negative-weight cycle detected" {
		t.Fatalf("expected error message 'negative-weight cycle detected', got: %v", err)
	}
}

func TestBellmanFordInvalidGraph(t *testing.T) {
	// Test with an empty graph
	g := graphs.NewWeightedGraph()
	_, err := FindPathWithBellmanFord(g, 0)
	if err == nil {
		t.Fatal("expected an error due to invalid graph, got nil")
	}

	// Test with a non-existent source node
	g.AddEdge(0, graphs.WeightedEdge{Node: 1, Weight: 1})
	_, err = FindPathWithBellmanFord(g, -1)
	if err == nil {
		t.Fatal("expected an error due to invalid source node, got nil")
	}
}
