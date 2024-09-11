package pathfinding

import (
	"dsa-go/structures/graphs"
	"fmt"
	"testing"
)

func TestFindPathWithDijkstra(t *testing.T) {

	// Create undirected graph
	g := graphs.NewWeightedGraph()

	g.AddEdge(0, graphs.WeightedEdge{Node: 1, Weight: 5})
	g.AddEdge(1, graphs.WeightedEdge{Node: 0, Weight: 5}) // Reverse edge

	g.AddEdge(0, graphs.WeightedEdge{Node: 2, Weight: 10})
	g.AddEdge(2, graphs.WeightedEdge{Node: 0, Weight: 10}) // Reverse edge

	g.AddEdge(1, graphs.WeightedEdge{Node: 3, Weight: 2})
	g.AddEdge(3, graphs.WeightedEdge{Node: 1, Weight: 2}) // Reverse edge

	g.AddEdge(2, graphs.WeightedEdge{Node: 3, Weight: 8})
	g.AddEdge(3, graphs.WeightedEdge{Node: 2, Weight: 8}) // Reverse edge

	g.AddEdge(3, graphs.WeightedEdge{Node: 4, Weight: 4})
	g.AddEdge(4, graphs.WeightedEdge{Node: 3, Weight: 4}) // Reverse edge

	g.AddEdge(1, graphs.WeightedEdge{Node: 4, Weight: 6}) // New edge from node 1 to node 4
	g.AddEdge(4, graphs.WeightedEdge{Node: 1, Weight: 6}) // Reverse edge for undirected graph

	g.AddEdge(2, graphs.WeightedEdge{Node: 4, Weight: 3}) // New edge from node 2 to node 4
	g.AddEdge(4, graphs.WeightedEdge{Node: 2, Weight: 3}) // Reverse edge for undirected graph

	// Expected distance table from node 0
	expectedDistances := []float64{
		0, 5, 10, 7, 11, // Distances from node 0 to nodes 0, 1, 2, 3, and 4
	}

	// Expected paths from node 0 to each node
	expectedPaths := [][]int{
		{0},       // Path from 0 to 0
		{0, 1},    // Path from 0 to 1
		{0, 2},    // Path from 0 to 2
		{0, 1, 3}, // Path from 0 to 3
		{0, 1, 4}, // Path from 0 to 4
	}

	// Print out the graph
	fmt.Println("Graph")
	graphs.PrintWeightedGraph(g)

	// Find the shortest path from node 0 using Dijkstra
	distanceTable, paths, err := FindPathWithDijkstra(g, 0)
	if err != nil {
		t.Fatalf("Error occurred during Dijkstra: %v", err)
	}

	// Validate the distance table
	for i, dist := range distanceTable {
		if dist != expectedDistances[i] {
			t.Errorf("Expected distance to node %d to be %.2f, but got %.2f", i, expectedDistances[i], dist)
		}
	}

	// Validate the paths
	for i, path := range paths {
		if len(path) != len(expectedPaths[i]) {
			t.Errorf("Expected path length to node %d to be %d, but got %d", i, len(expectedPaths[i]), len(path))
		}
		for j := range path {
			if path[j] != expectedPaths[i][j] {
				t.Errorf("Expected path to node %d to be %v, but got %v", i, expectedPaths[i], path)
				break
			}
		}
	}

}
