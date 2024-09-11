package pathfinding

import (
	"dsa-go/structures/graphs"
	"dsa-go/structures/queue"
	"errors"
	"math"
)

// Heuristic function
func heuristic(node, goal int) float64 {
	return math.Abs(float64(node - goal)) // Simple difference
}

// Function to reconstruct the path from the parentTable
func reconstructAStarPath(parentTable []int, currentNode int) []int {
	var path []int
	for currentNode != -1 {
		path = append([]int{currentNode}, path...) // Prepend current node to the path
		currentNode = parentTable[currentNode]
	}
	return path
}

func FindPathWithAStar(g *graphs.WeightedGraph, src, dest int) ([]int, error) {

	// Input validation
	if src < 0 || len(g.Nodes) == 0 || g.Nodes[src] == nil {
		return nil, errors.New("invalid graph or source node")
	}

	// Initialize required tables
	V := len(g.Nodes)

	// g-score - distance from start node to each node, initialize to infinity
	gScore := make([]float64, V)
	for i := 0; i < V; i++ {
		gScore[i] = math.Inf(1) // Infinity
	}
	gScore[src] = 0 // Distance to the source node is 0

	// f-score - estimated total cost (g + h), initialize to infinity
	fScore := make([]float64, V)
	for i := 0; i < V; i++ {
		fScore[i] = math.Inf(1) // Infinity
	}
	fScore[src] = heuristic(src, dest) // Heuristic from source to destination

	// parentTable - to track the best path, initialize with -1 (no parent)
	parentTable := make([]int, V)
	for i := 0; i < V; i++ {
		parentTable[i] = -1
	}

	// openSet - priority queue to store nodes to be evaluated
	openList := queue.CreatePriorityQueue[int]()
	openList.EnqueuePriority(src, int(fScore[src]))

	// closedSet - to store evaluated nodes
	closedList := make(map[int]bool)

	for !openList.IsEmpty() {

		// Extract node with minimum distance
		currentNode, err := openList.DequeuePriority()
		if err != nil {
			return nil, errors.New("error dequeuing from priority queue")
		}

		// Check if the current node is the goal
		if currentNode == dest {
			// Goal reached, reconstruct the path
			path := reconstructAStarPath(parentTable, currentNode)
			return path, nil
		}

		// Add current node to the closed list as it's now evaluated
		closedList[currentNode] = true

		// Iterate through neighbours
		for _, edge := range g.Nodes[currentNode].GetItems() {

			neighbor := edge.Node
			distance := edge.Weight

			// If the neighbor is in the closed list, skip it
			if closedList[neighbor] == true {
				continue
			}

			// Calculate the tentative g_score of the neighbor
			tentativeGScore := gScore[currentNode] + float64(distance)

			if tentativeGScore < gScore[neighbor] {
				// Update g-score and f-score
				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = gScore[neighbor] + heuristic(neighbor, dest)

				// Set the current node as the parent of the neighbor
				parentTable[neighbor] = currentNode

				// Add the neighbor to the open list with its f-score
				openList.EnqueuePriority(neighbor, int(fScore[neighbor]))
			}

		}
	}

	// If the open set is empty, no path was found
	return nil, errors.New("no path exists from source to destination")
}
