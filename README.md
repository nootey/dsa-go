# Go Data structures and algorithms

This repository contains implementations of various data structures and algorithms using GoLang.
The goal is to implement as many as possible while learning and mastering GoLang.

Each implementation is accompanied by unit tests to ensure correctness.

To run the tests, you can use the following command.

```bash
go run RunTests.go
```

To run tests in a single directory, use the command.

```bash
# To run tests on the arrays directory
go run RunTests.go ./structures/arrays
```

## Data Structures

  - [x] [Arrays (One-dimensional)](./structures/arrays/1DArray.go)
  - [x] [Arrays (Two-dimensional)](./structures/arrays/2DArray.go)
  - [x] [Singly Linked List](./structures/linked-lists/SinglyLinkedList.go)
  - [x] [Doubly Linked List](./structures/linked-lists/DoublyLinkedList.go)
  - [x] [Stack](./structures/stacks/Stack.go)
  - [x] [Simple Queue](./structures/queue/SimpleQueue.go)
  - [x] [Priority Queue](./structures/queue/PriorityQueue.go)
  - [x] [Heap Priority Queue](./structures/queue/PriorityQueueHeap.go)
  - [x] [Hash Table](./structures/hash-tables/HashTable.go)
  - [x] [Binary Search Tree](./structures/trees/BinarySearchTree.go)
  - [x] [Heap](./structures/heap/Heap.go)

 ### Graph representation

  - [x] [Adjacency Matrix](./structures/graphs/AdjacencyMatrix.go)
  - [x] [Adjacency List](./structures/graphs/AdjacencyList.go)

## Algorithms

### Sorting

  - [x] [Bubble sort](./algorithms/sorting/BubbleSort.go)
  - [x] [Selection sort](./algorithms/sorting/SelectionSort.go)
  - [x] [Insertion sort](./algorithms/sorting/InsertionSort.go)
  - [x] [Merge sort](./algorithms/sorting/MergeSort.go)
  - [x] [Quick sort](./algorithms/sorting/QuickSort.go)
  - [ ] TODO: [Heap sort](./algorithms/sorting/HeapSort.go)

### Searching

  - [x] [Linear Search](./algorithms/searching/LinearSearch.go)
  - [x] [Binary Search](./algorithms/searching/BinarySearch.go)

### Graphing

#### Basic traversal

  - [x] [Breadth-First Search](./algorithms/graphing/basic-traversal/BreadthFirstSearch.go)
  - [x] [Depth-First Search](./algorithms/graphing/basic-traversal/DepthFirstSearch.go)

#### Pathfinding

  - [x] [Dijkstra](./algorithms/graphing/pathfinding/Dijkstra.go)
  -     TODO: Implement Heap Priority Queue to Dijkstra
  - [x] [Bellman-Ford](./algorithms/graphing/pathfinding/BellmanFord.go)
  - [x] [Floyd-Warshall](./algorithms/graphing/pathfinding/FloydWarshall.go)
  - [x] [A* (A-Star)](./algorithms/graphing/pathfinding/AStar.go)
 