package pathfinding

import (
	"dsa-go/structures/graphs"
	"fmt"
)

const INF int = 1 << 30

func FindPathWithFloydWarshall(g *graphs.AdjacencyMatrixGraph) ([][]int, [][]int) {

	// Number of vertices in the graph
	v := len(g.Matrix)

	distanceTable := make([][]int, v)
	predecessorTable := make([][]int, v)

	// Initialize the distanceTable with graph weights
	for i := 0; i < v; i++ {
		distanceTable[i] = make([]int, v)
		predecessorTable[i] = make([]int, v)
		for j := 0; j < v; j++ {
			if i == j {
				distanceTable[i][j] = 0
				predecessorTable[i][j] = -1 // No predecessor for the same node
			} else if g.Matrix[i][j] != 0 {
				distanceTable[i][j] = g.Matrix[i][j]
				predecessorTable[i][j] = i // Initialize with the start node
			} else {
				distanceTable[i][j] = INF
				predecessorTable[i][j] = -1 // No path yet
			}
		}
	}

	// i: Represents the source vertex - the starting point of the path
	// j: Represents the destination vertex -  endpoint of the path
	// k: Represents an intermediate vertex - the vertex through which determines whether a shorter path from i to j exists compared to the current known shortest path.

	// Find the shortest path
	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				// Ensure there is a valid path between i and k, && k and j
				if distanceTable[i][k] < INF && distanceTable[k][j] < INF && distanceTable[i][j] > distanceTable[i][k]+distanceTable[k][j] {
					distanceTable[i][j] = distanceTable[i][k] + distanceTable[k][j]
					predecessorTable[i][j] = predecessorTable[k][j] // Update the path
				}
			}
		}
	}

	printDistanceTable(distanceTable, v)
	printPredecessorTable(predecessorTable, v)

	return distanceTable, predecessorTable
}

func printDistanceTable(distanceTable [][]int, v int) {
	// Print distanceTable nicely
	fmt.Println("Distance Table:")
	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			if distanceTable[i][j] == INF {
				fmt.Print("INF ")
			} else {
				fmt.Printf("%d ", distanceTable[i][j])
			}
		}
		fmt.Println()
	}
}

func printPredecessorTable(predecessorTable [][]int, v int) {
	fmt.Println("Predecessor Table:")
	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			if predecessorTable[i][j] == -1 {
				fmt.Print("NIL ")
			} else {
				fmt.Printf("%d ", predecessorTable[i][j])
			}
		}
		fmt.Println()
	}
}

func ReconstructFloydPath(predecessorTable [][]int, start, end int) []int {
	if predecessorTable[start][end] == -1 {
		return nil // No path exists
	}

	var path []int
	for end != -1 {
		path = append([]int{end}, path...)
		end = predecessorTable[start][end]
	}
	return path
}
