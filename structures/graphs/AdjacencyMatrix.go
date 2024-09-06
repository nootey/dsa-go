package graphs

import "fmt"

type Node struct {
	data string
}

type AdjacencyMatrixGraph struct {
	matrix [][]int
	nodes  []Node
}

// CreateNewGraph creates a new graph with the specified size (number of vertices)
func CreateNewGraph(size int) *AdjacencyMatrixGraph {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size) // Initialize each row
	}
	return &AdjacencyMatrixGraph{matrix: matrix}
}

// AddNode adds a new node to the graph
func (g *AdjacencyMatrixGraph) AddNode(data string) {
	node := Node{data: data}
	g.nodes = append(g.nodes, node)
}

// AddEdge adds a directed edge between two vertices
func (g *AdjacencyMatrixGraph) AddEdge(src int, dest int) {
	// Ensure that src and dest are within the bounds of the graph size
	if src < len(g.matrix) && dest < len(g.matrix) {
		g.matrix[src][dest] = 1
	}
}

// CheckEdge checks if there is an edge between two vertices
func (g *AdjacencyMatrixGraph) CheckEdge(src int, dest int) bool {
	// Ensure that src and dest are within the bounds of the graph size
	if src < len(g.matrix) && dest < len(g.matrix) {
		return g.matrix[src][dest] == 1
	}
	return false
}

// PrintMatrix prints the adjacency matrix of the graph
func (g *AdjacencyMatrixGraph) PrintMatrix() {
	fmt.Println("Adjacency Matrix:")
	for _, row := range g.matrix {
		fmt.Println(row)
	}
}
