package basic_traversal

import (
	"dsa-go/structures/graphs"
	"dsa-go/structures/stacks"
	"errors"
	"fmt"
)

func DepthFirstSearchRecursive(g *graphs.AdjacencyMatrixGraph, src int) error {

	if src < 0 {
		return errors.New("provided node index out of range")
	}

	if len(g.Nodes) == 0 {
		return errors.New("graph has no nodes")
	}

	if g.Matrix[src] == nil {
		return errors.New("wanted node does not exist")
	}

	visited := make([]bool, len(g.Matrix))
	IsVisited(g, src, visited)

	return nil
}

func IsVisited(g *graphs.AdjacencyMatrixGraph, src int, visited []bool) {
	if visited[src] {
		return
	} else {
		// Optional print out progress
		fmt.Println("Visited: " + g.Nodes[src].Data)
		visited[src] = true
	}

	for i := range g.Matrix[src] {
		if g.Matrix[src][i] == 1 {
			IsVisited(g, i, visited)
		}
	}

	return
}

func DepthFirstSearchStack(g *graphs.AdjacencyMatrixGraph, src int) error {

	if src < 0 {
		return errors.New("provided node index out of range")
	}

	if len(g.Nodes) == 0 {
		return errors.New("graph has no nodes")
	}

	if g.Matrix[src] == nil {
		return errors.New("wanted node does not exist")
	}

	visited := make([]bool, len(g.Matrix))
	stack := stacks.CreateStack[int]()

	// Push the starting node onto the stack
	stack.Push(src)

	for !stack.IsEmpty() {
		// Pop a node from the stack
		current, _ := stack.Pop()

		// Check if the current node has already been visited
		if visited[current] {
			continue
		}

		// Mark the node as visited
		fmt.Println("Visited: " + g.Nodes[current].Data)
		visited[current] = true

		// Push all unvisited neighbors onto the stack
		for i := range g.Matrix[current] {
			if g.Matrix[current][i] == 1 && !visited[i] {
				stack.Push(i)
			}
		}
	}

	return nil
}
