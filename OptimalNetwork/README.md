# Main Code

## Node and PriorityQueue Structures:

- The `Node` struct represents a router and includes the accumulated latency and a boolean indicating whether compression has been used.
- The `PriorityQueue` struct is a min-heap of `Node` pointers to prioritize nodes with the smallest latency.

## Dijkstra's Algorithm:

- Initialize a priority queue and push the source node with zero latency.
- Use a map to track the minimum latency to each node.
- Pop nodes from the priority queue, update latencies for neighboring nodes, and push them back into the queue.
- For nodes with compression capability, push an additional node into the queue with halved latency if compression has not been used yet.