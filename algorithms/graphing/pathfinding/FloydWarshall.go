package pathfinding

import (
	"dsa-go/structures/graphs"
	"fmt"
)

const INF int = 1 << 30

func FindPathWithFloydWarshall(g *graphs.AdjacencyMatrixGraph) [][]int {

	// Number of vertices in the graph
	v := len(g.Matrix)

	distanceTable := make([][]int, v)

	// Initialize the distanceTable with graph weights
	for i := 0; i < v; i++ {
		distanceTable[i] = make([]int, v)
		for j := 0; j < v; j++ {
			if i == j {
				distanceTable[i][j] = 0
			} else if g.Matrix[i][j] != 0 {
				distanceTable[i][j] = g.Matrix[i][j]
			} else {
				distanceTable[i][j] = INF
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
				}
			}
		}
	}

	printDistanceTable(distanceTable, v)

	return distanceTable
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
