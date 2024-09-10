package pathfinding

import (
	"dsa-go/structures/graphs"
	"fmt"
	"log"
	"testing"
)

func TestFindPathWithFloydWarshall(t *testing.T) {
	// Create a new adjacency matrix graph with 5 nodes
	g, err := graphs.CreateNewGraph(5)
	if err != nil {
		log.Fatal(err)
	}

	// Add edges to the graph
	g.AddEdge(0, 1, 3) // Edge from node 0 to node 1 with weight 3
	g.AddEdge(0, 3, 7) // Edge from node 0 to node 3 with weight 7
	g.AddEdge(1, 2, 1) // Edge from node 1 to node 2 with weight 1
	g.AddEdge(1, 3, 5) // Edge from node 1 to node 3 with weight 5
	g.AddEdge(2, 3, 2) // Edge from node 2 to node 3 with weight 2
	g.AddEdge(3, 4, 1) // Edge from node 3 to node 4 with weight 1
	g.AddEdge(4, 0, 2) // Edge from node 4 to node 0 with weight 2

	// Expected output after running Floyd-Warshall
	expectedDistances := [][]int{
		{0, 3, 4, 6, 7}, // From node 0
		{6, 0, 1, 3, 4}, // From node 1
		{5, 8, 0, 2, 3}, // From node 2
		{3, 6, 7, 0, 1}, // From node 3
		{2, 5, 6, 8, 0}, // From node 4
	}

	fmt.Println("\n")
	// Run Floyd-Warshall algorithm
	distanceTable, predecessorTable := FindPathWithFloydWarshall(g)

	// Validate the distance table against the expected output
	for i := range distanceTable {
		for j := range distanceTable[i] {
			if distanceTable[i][j] != expectedDistances[i][j] {
				t.Errorf("Distance mismatch at (%d, %d): expected %d, got %d",
					i, j, expectedDistances[i][j], distanceTable[i][j])
			}
		}
	}

	// Example: Reconstruct the path from node 0 to node 4
	expectedPath := []int{0, 1, 2, 3, 4}
	src := 0
	dst := 4
	path := ReconstructFloydPath(predecessorTable, src, dst)
	if len(path) != len(expectedPath) {
		t.Errorf("Path length mismatch: expected %v, got %v", expectedPath, path)
	}
	for i := range path {
		if path[i] != expectedPath[i] {
			t.Errorf("Path mismatch at position %d: expected %d, got %d", i, expectedPath[i], path[i])
		}
	}

	fmt.Printf("Path for node from %d to %d: %d", src, dst, path)

}
