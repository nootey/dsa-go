package basic_traversal

import (
	"dsa-go/structures/graphs"
	queues "dsa-go/structures/queue"
	"errors"
	"fmt"
)

func BreadthFirstSearch(g *graphs.AdjacencyMatrixGraph, src int) error {

	if src < 0 {
		return errors.New("provided node index out of range")
	}

	if len(g.Nodes) == 0 {
		return errors.New("graph has no nodes")
	}

	if g.Matrix[src] == nil {
		return errors.New("wanted node does not exist")
	}

	queue := queues.CreateSimpleQueue[int](len(g.Matrix))
	visited := make([]bool, len(g.Matrix))

	err := queue.EnqueueSimple(src)
	if err != nil {
		return err
	}

	visited[src] = true

	for !queue.IsEmpty() {

		src, err = queue.DequeueSimple()
		if err != nil {
			return err
		}
		fmt.Println("Visited: " + g.Nodes[src].Data)

		for i := range g.Matrix[src] {
			if g.Matrix[src][i] == 1 && !visited[i] {
				err := queue.EnqueueSimple(i)
				if err != nil {
					return err
				}
				visited[i] = true
			}
		}

	}

	return nil
}
