package ng

import (
	"container/heap"
	"fmt"
	"math"
)

// Node represents a tile in the pathfinding graph
type Node struct {
	X, Y     int
	F, G, H  float64 // costs for A* algorithm
	Parent   *Node
	MoveCost MoveCost
}

// PriorityQueue is the open set of nodes to explore
// implements heap.Interface for A* pathfinding
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F // Lower F value has higher priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[0]     // Get the root element (lowest F value)
	old[0] = old[n-1]  // Move the last element to the root
	*pq = old[0 : n-1] // Remove the last element
	heap.Fix(pq, 0)    // Restore heap property after the swap
	return item
}

// FindPath finds the optimal path between two tiles in a region using A* algorithm
func (region *Region) FindPath(startX, startY, endX, endY int) []Position {
	// Initialize open and closed sets
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	closedSet := make(map[string]bool)

	// Create start and end nodes
	startNode := &Node{
		X:        startX,
		Y:        startY,
		MoveCost: region.Tiles[startX][startY].MoveCost,
	}
	endNode := &Node{
		X: endX,
		Y: endY,
	}

	// Initialize start node
	startNode.G = 0
	startNode.H = heuristic(startNode, endNode)
	startNode.F = startNode.G + startNode.H

	// Add start node to open set
	heap.Push(openSet, startNode)

	// Main A* loop
	for openSet.Len() > 0 {
		// Get node with lowest F cost
		current := heap.Pop(openSet).(*Node)

		// Check if we reached the goal
		if current.X == endNode.X && current.Y == endNode.Y {
			return reconstructPath(current)
		}

		// Add current node to closed set
		closedSet[getNodeKey(current)] = true

		// Check neighbors
		for _, neighbor := range getNeighbors(region, current) {
			// Skip if neighbor is in closed set
			if closedSet[getNodeKey(neighbor)] {
				continue
			}

			// Calculate new G cost
			newG := current.G + float64(neighbor.MoveCost)

			// Skip if this path is worse than previous
			if !containsNode(openSet, neighbor) {
				heap.Push(openSet, neighbor)
			} else if newG >= neighbor.G {
				continue
			}

			// This path is the best so far, record it
			neighbor.Parent = current
			neighbor.G = newG
			neighbor.H = heuristic(neighbor, endNode)
			neighbor.F = neighbor.G + neighbor.H
		}
	}

	// No path found
	return nil
}

// heuristic calculates the diagonal distance between two nodes
func heuristic(a, b *Node) float64 {
	dx := math.Abs(float64(a.X - b.X))
	dy := math.Abs(float64(a.Y - b.Y))
	// Using diagonal distance formula: max(dx, dy) + (sqrt2 - 1) * min(dx, dy)
	return math.Max(dx, dy) + 0.414*math.Min(dx, dy)
}

// getNeighbors returns all valid neighboring nodes including diagonals
func getNeighbors(region *Region, node *Node) []*Node {
	neighbors := make([]*Node, 0, 8) // Increased capacity for 8 directions
	directions := [][2]int{
		{0, 1},   // up
		{1, 1},   // up-right
		{1, 0},   // right
		{1, -1},  // down-right
		{0, -1},  // down
		{-1, -1}, // down-left
		{-1, 0},  // left
		{-1, 1},  // up-left
	}

	for _, dir := range directions {
		newX, newY := node.X+dir[0], node.Y+dir[1]

		// Check bounds
		if newX < 0 || newX >= len(region.Tiles) || newY < 0 || newY >= len(region.Tiles[0]) {
			continue
		}

		// Check if tile is passable
		moveCost := region.Tiles[newX][newY].MoveCost
		if moveCost == ImpassableCost {
			continue
		}

		// For diagonal movement, check if both adjacent tiles are passable

		// if dir[0] != 0 && dir[1] != 0 {
		// 	// Check the two adjacent tiles that form the diagonal
		// 	adjX1, adjY1 := node.X+dir[0], node.Y
		// 	adjX2, adjY2 := node.X, node.Y+dir[1]

		// 	if region.Tiles[adjX1][adjY1].MoveCost == ImpassableCost ||
		// 		region.Tiles[adjX2][adjY2].MoveCost == ImpassableCost {
		// 		continue
		// 	}
		// }

		// Calculate diagonal movement cost (sqrt of 2)
		if dir[0] != 0 && dir[1] != 0 {
			moveCost = moveCost * 1.414
		}

		neighbors = append(neighbors, &Node{
			X:        newX,
			Y:        newY,
			MoveCost: moveCost,
		})
	}

	return neighbors
}

// getNodeKey returns a unique key for a node
func getNodeKey(node *Node) string {
	return fmt.Sprintf("%d,%d", node.X, node.Y)
}

// containsNode checks if a node exists in the priority queue
func containsNode(pq *PriorityQueue, node *Node) bool {
	for _, n := range *pq {
		if n.X == node.X && n.Y == node.Y {
			return true
		}
	}
	return false
}

// reconstructPath builds the path from end node to start node, excluding the initial position
func reconstructPath(endNode *Node) []Position {
	path := make([]Position, 0)
	current := endNode

	for current != nil && current.Parent != nil {
		path = append([]Position{{X: current.X, Y: current.Y}}, path...)
		current = current.Parent
	}

	return path
}
