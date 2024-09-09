package graphs

import (
	"errors"
	"fmt"
)

type Node struct {
	Data string
}

type AdjacencyMatrixGraph struct {
	Matrix [][]int
	Nodes  []Node
}

// CreateNewGraph creates a new graph with the specified size (number of vertices)
func CreateNewGraph(size int) (*AdjacencyMatrixGraph, error) {
	if size < 0 {
		return nil, errors.New("invalid size")
	}
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size) // Initialize each row
	}
	return &AdjacencyMatrixGraph{Matrix: matrix}, nil
}

// AddNode adds a new node to the graph
func (g *AdjacencyMatrixGraph) AddNode(data string) {
	node := Node{Data: data}
	g.Nodes = append(g.Nodes, node)
}

// AddEdge adds a directed edge between two vertices
func (g *AdjacencyMatrixGraph) AddEdge(src int, dest int, w ...int) {

	weight := 1
	if len(w) > 0 {
		weight = w[0]
	}

	// Ensure that src and dest are within the bounds of the graph size
	if src < len(g.Matrix) && dest < len(g.Matrix) {
		g.Matrix[src][dest] = weight
	}
}

// CheckEdge checks if there is an edge between two vertices
func (g *AdjacencyMatrixGraph) CheckEdge(src int, dest int) bool {
	// Ensure that src and dest are within the bounds of the graph size
	if src < len(g.Matrix) && dest < len(g.Matrix) {
		return g.Matrix[src][dest] == 1
	}
	return false
}

// PrintMatrix prints the adjacency matrix of the graph
func (g *AdjacencyMatrixGraph) PrintMatrix() {
	fmt.Println("Adjacency Matrix:")
	for _, row := range g.Matrix {
		fmt.Println(row)
	}
}
