package pathfinding

import (
	"dsa-go/structures/graphs"
	"dsa-go/structures/queue"
	"errors"
	"fmt"
	"math"
)

//TODO: Currently uses a non-heap priority queue

func FindPathWithDijkstra(g *graphs.WeightedGraph, src int) ([]float64, [][]int, error) {

	// Input validation
	if src < 0 || len(g.Nodes) == 0 || g.Nodes[src] == nil {
		return nil, nil, errors.New("invalid graph or source node")
	}

	// Initialize required tables
	distanceTable := make([]float64, len(g.Nodes))
	previousNodeTable := make([]*int, len(g.Nodes))
	visitedNodesTable := make([]bool, len(g.Nodes))

	for i := range distanceTable {
		distanceTable[i] = math.Inf(1) // Initialize with infinity
		previousNodeTable[i] = nil
	}
	distanceTable[src] = 0

	// Priority Queue of nodes with their distances
	pq := queue.CreatePriorityQueue[int]() // Create a priority queue for nodes
	pq.EnqueuePriority(src, 0)             // Add the source node to the queue with distance 0

	for !pq.IsEmpty() {
		// Extract node with minimum distance
		u, err := pq.DequeuePriority()
		if err != nil {
			return nil, nil, errors.New("error dequeuing from priority queue")
		}

		// If this node is already visited, skip it
		if visitedNodesTable[u] {
			continue
		}

		visitedNodesTable[u] = true

		// Verify node u is within bounds and has neighbors
		if u < 0 || u >= len(g.Nodes) || g.Nodes[u] == nil {
			continue
		}

		neighbors := g.Nodes[u].GetItems()

		// For each neighbor of u, check if a shorter path to the neighbor v is found through u
		for _, edge := range neighbors {
			v := edge.Node
			weight := edge.Weight

			// Ensure the destination node v is within bounds
			if v >= 0 && v < len(distanceTable) {
				// Relax the edge (u, v)
				if distanceTable[u]+float64(weight) < distanceTable[v] {

					distanceTable[v] = distanceTable[u] + float64(weight)
					previousNodeTable[v] = &u

					// Add the neighbor to the priority queue with updated distance
					pq.EnqueuePriority(v, int(distanceTable[v]))
				}
			}
		}

	}

	// Reconstruct paths for each node
	paths := make([][]int, len(g.Nodes))
	for i := range paths {
		paths[i] = reconstructPath(previousNodeTable, src, i)
	}

	fmt.Println("Distance Table:", distanceTable)
	fmt.Println("Paths:", paths)

	return distanceTable, paths, nil
}

// reconstructPath reconstructs the path from source to destination
func reconstructPath(previousNodeTable []*int, src, dest int) []int {
	var path []int

	// Backtrack from destination to source using the previousNodeTable
	for current := &dest; current != nil; current = previousNodeTable[*current] {
		path = append([]int{*current}, path...) // Add the current node to the front of the path
		if *current == src {
			break
		}
	}

	// If the destination is unreachable, return an empty path
	if len(path) == 0 || path[0] != src {
		return []int{}
	}

	return path
}

func main() {

}
