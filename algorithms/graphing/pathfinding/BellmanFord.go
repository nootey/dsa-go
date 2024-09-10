package pathfinding

import (
	"dsa-go/structures/graphs"
	"errors"
	"fmt"
	"math"
)

func FindPathWithBellmanFord(g *graphs.WeightedGraph, src int) ([]float64, error) {
	// Input validation
	if src < 0 || len(g.Nodes) == 0 || g.Nodes[src] == nil {
		return nil, errors.New("invalid graph or source node")
	}

	// Initialize required tables
	V := len(g.Nodes)
	distanceTable := make([]float64, V)
	previousNodeTable := make([]*int, V)

	for i := range distanceTable {
		distanceTable[i] = math.Inf(1) // Initialize with infinity
		previousNodeTable[i] = nil
	}
	distanceTable[src] = 0

	// Step 1: Relax edges V-1 times
	for i := 0; i < V-1; i++ {
		for u := range g.Nodes {
			edges := g.Nodes[u].GetItems()
			for _, edge := range edges {
				v := edge.Node
				weight := edge.Weight

				// Relax the edge (u -> v)
				if distanceTable[u]+float64(weight) < distanceTable[v] {
					distanceTable[v] = distanceTable[u] + float64(weight)
					previousNodeTable[v] = &u
				}
			}
		}
	}

	// Step 2: Check for negative-weight cycles
	for u := range g.Nodes {
		edges := g.Nodes[u].GetItems()
		for _, edge := range edges {
			v := edge.Node
			weight := edge.Weight

			// If an edge can still be relaxed, there is a negative-weight cycle
			if distanceTable[u]+float64(weight) < distanceTable[v] {
				return nil, errors.New("negative-weight cycle detected")
			}
		}
	}

	fmt.Println("Distance Table:", distanceTable)

	return distanceTable, nil
}
