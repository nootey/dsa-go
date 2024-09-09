package graphs

import (
	linkedList "dsa-go/structures/linked-lists"
	"fmt"
)

type WeightedEdge struct {
	Weight int
	Node   int
}

type WeightedGraph struct {
	Nodes []*linkedList.SinglyLinkedList[WeightedEdge]
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		Nodes: []*linkedList.SinglyLinkedList[WeightedEdge]{},
	}
}

func (g *WeightedGraph) ensureCapacity(node int) {
	if node >= len(g.Nodes) {
		// Resize the slice to accommodate the new node
		newSize := node + 1
		newNodes := make([]*linkedList.SinglyLinkedList[WeightedEdge], newSize)
		copy(newNodes, g.Nodes)
		g.Nodes = newNodes
	}
}

func (g *WeightedGraph) AddEdge(node int, child WeightedEdge) {
	// Ensure the source node has capacity
	g.ensureCapacity(node)

	// Ensure the destination node exists, even if it has no outgoing edges yet
	g.ensureCapacity(child.Node)

	// Initialize the adjacency list for the source node if needed
	if g.Nodes[node] == nil {
		g.Nodes[node] = linkedList.NewSinglyLinkedList[WeightedEdge]()
	}

	// Add the edge to the source node's adjacency list
	err := g.Nodes[node].AddItem(child)
	if err != nil {
		fmt.Printf("Error adding edge: %v\n", err)
	}
}

func PrintWeightedGraph(g *WeightedGraph) {
	for nodeIndex, adjacencyList := range g.Nodes {
		if adjacencyList == nil {
			continue
		}

		edges := adjacencyList.GetItems()
		if len(edges) == 0 {
			continue
		}

		fmt.Printf("Node %d has edges:\n", nodeIndex)
		for _, edge := range edges {
			fmt.Printf("  -> Node %d (Weight: %d)\n", edge.Node, edge.Weight)
		}
	}
}
