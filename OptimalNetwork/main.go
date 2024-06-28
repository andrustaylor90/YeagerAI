package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a router and the accumulated latency to reach this node.
type Node struct {
	id         string
	latency    int
	compressed bool
}

// PriorityQueue is a priority queue of Nodes, implementing heap.Interface.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].latency < pq[j].latency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// Dijkstra's algorithm with compression capability.
func find_minimum_latency_path(graph map[string][][2]interface{}, compression_nodes []string, source, destination string) int {
	compressionSet := make(map[string]bool)
	for _, node := range compression_nodes {
		compressionSet[node] = true
	}

	// Priority queue to select the node with the smallest latency.
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{id: source, latency: 0, compressed: false})

	// Maps to track the minimum latency to each node.
	minLatency := make(map[string]int)
	for node := range graph {
		minLatency[node] = math.MaxInt32
	}
	minLatency[source] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		if current.id == destination {
			return current.latency
		}

		for _, edge := range graph[current.id] {
			neighbor := edge[0].(string)
			weight := edge[1].(int)
			newLatency := current.latency + weight

			if newLatency < minLatency[neighbor] {
				minLatency[neighbor] = newLatency
				heap.Push(pq, &Node{id: neighbor, latency: newLatency, compressed: current.compressed})
			}

			if !current.compressed && compressionSet[current.id] {
				compressedLatency := current.latency + weight/2
				if compressedLatency < minLatency[neighbor] {
					minLatency[neighbor] = compressedLatency
					heap.Push(pq, &Node{id: neighbor, latency: compressedLatency, compressed: true})
				}
			}
		}
	}

	return -1 // Destination not reachable
}

func main() {
	graph := map[string][][2]interface{}{
		"A": {{"B", 10}, {"C", 20}},
		"B": {{"D", 15}},
		"C": {{"D", 30}},
		"D": {},
	}
	compression_nodes := []string{"B", "C"}
	source := "A"
	destination := "D"
	min_latency := find_minimum_latency_path(graph, compression_nodes, source, destination)
	fmt.Printf("Minimum total latency: %d\n", min_latency)
}
