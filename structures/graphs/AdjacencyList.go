package graphs

import (
	linkedList "dsa-go/structures/linked-lists"
	"log"
)

type AdjacencyListGraph struct {
	nodes map[int]*linkedList.SinglyLinkedList[int]
}

// CreateNewGraphList creates a new graph with the specified size (number of vertices)
func CreateNewGraphList() *AdjacencyListGraph {
	return &AdjacencyListGraph{nodes: map[int]*linkedList.SinglyLinkedList[int]{}}
}

// AddEdge adds a directed edge between two vertices
func (g *AdjacencyListGraph) AddEdge(src int, dst int) {

	// Makes sure to add a node, if list is empty
	if g.nodes[src] == nil {
		g.nodes[src] = linkedList.NewSinglyLinkedList[int]()
	}

	err := g.nodes[src].AddItem(dst)
	if err != nil {
		log.Fatal(err)
	}
}

// CheckEdge checks if the destination vertex is in the linked list of the source vertex
func (g *AdjacencyListGraph) CheckEdge(src int, dest int) bool {

	// Ensure src exists in the graph
	if _, exists := g.nodes[src]; !exists {
		return false
	}

	node, _ := g.nodes[src].FindItem(dest)
	return node != nil
}

// PrintList prints the adjacency matrix of the graph
func (g *AdjacencyListGraph) PrintList() {
	for vertex, list := range g.nodes {

		log.Printf("Vertex %d:", vertex)

		// Iterate through the linked list and print adjacent vertices
		current := list.Head
		if current == nil {
			log.Println("  No edges")
			continue
		}
		for current != nil {
			log.Printf("  -> %d", current.Value)
			current = current.Next
		}
	}
}
