package graphs

import (
	linkedList "dsa-go/structures/linked-lists"
	"log"
)

type AdjacencyListGraph struct {
	Nodes map[int]*linkedList.SinglyLinkedList[int]
}

// CreateNewGraphList creates a new graph with the specified size (number of vertices)
func CreateNewGraphList() *AdjacencyListGraph {
	return &AdjacencyListGraph{Nodes: map[int]*linkedList.SinglyLinkedList[int]{}}
}

// AddEdge adds a directed edge between two vertices
func (g *AdjacencyListGraph) AddEdge(src int, dst int) {

	// Makes sure to add a node, if list is empty
	if g.Nodes[src] == nil {
		g.Nodes[src] = linkedList.NewSinglyLinkedList[int]()
	}

	err := g.Nodes[src].AddItem(dst)
	if err != nil {
		log.Fatal(err)
	}
}

// CheckEdge checks if the destination vertex is in the linked list of the source vertex
func (g *AdjacencyListGraph) CheckEdge(src int, dest int) bool {

	// Ensure src exists in the graph
	if _, exists := g.Nodes[src]; !exists {
		return false
	}

	node, _ := g.Nodes[src].FindItem(dest)
	return node != nil
}

// PrintList prints the adjacency matrix of the graph
func (g *AdjacencyListGraph) PrintList() {
	for vertex, list := range g.Nodes {

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
