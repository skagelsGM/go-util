package graph

import (
	"fmt"

	"github.com/skagelsGM/util/list"
)

// Graph is an implementation of the graph data structure, consisting of a finite set of vertices
// and edges. Edges are represents using an array of adjacent node list. Graph is undirected.
type Graph struct {
	Vertices int
	AdjNodes []list.LinkedList
}

// CreateGraph creates a graph with v number of vertices (nodes)
func CreateGraph(v int) *Graph {
	var g = new(Graph)
	g.Vertices = v
	g.AdjNodes = make([]list.LinkedList, v)

	//initialize adjacent node lists
	for i := 0; i < v; i++ {
		var list = new(list.LinkedList)
		g.AdjNodes[i] = *list
	}

	return g
}

// AddEdge adds an edge to Graph between nodes src and dest.
// Graph edges are undirected.
func (g *Graph) AddEdge(src, dest int) {
	g.AdjNodes[src].Add(dest)
	// Add the edge from dest to src, too, since Graph is undirected.
	g.AdjNodes[dest].Add(src)
	fmt.Printf("Added Edge: %d-%d\n", src, dest)
}

// IsAdjacent returns true if src and dest are adjacent nodes
func (g *Graph) IsAdjacent(src, dest int) bool {
	edge := g.AdjNodes[src].Head
	if edge == nil {
		return false
	}

	for {
		if edge.Value == dest {
			return true
		}

		if edge.Next == nil {
			break
		}

		edge = edge.Next
	}

	return false
}

// BreadthFirstSearch traverses the graph from the given start node using BFS.
// Returns the path traversed as a linked list.
func (g *Graph) BreadthFirstSearch(start int) *list.LinkedList {
	fmt.Printf("[Graph.BreadthFirstSearch] Performing BFS on graph, starting at node %d\n", start)
	// queue is the ordered list of nodes to traverse
	queue := new(list.LinkedList)
	// path is final list of visited nodes
	path := new(list.LinkedList)
	// track visited nodes
	visited := make(map[int]bool)
	queue.Add(start)

	// grab next node from queue
	for queue.Head != nil {
		node := queue.Remove()

		if visited[node.Value] {
			continue
		}

		path.Add(node.Value)
		visited[node.Value] = true
		fmt.Printf("Visited Node: %d\n", node.Value)

		// Enqueue adjacent nodes for visiting:
		// - Don't fall into the trap of adding the next adjacent node. Enqueue a new node to keep
		// adjacent linked lists and queue separate.
		for next := g.AdjNodes[node.Value].Head; next != nil; next = next.Next {
			queue.Add(next.Value)
			fmt.Printf("(added adjacent node %d to the queue)\n", next.Value)
		}
	}

	return path
}

// DepthFirstSearch traverses the graph from the given start node using DFS.
// Returns the path traversed as a linked list.
func (g *Graph) DepthFirstSearch(start int) *list.LinkedList {
	fmt.Printf("[Graph.DepthFirstSearch] Performing DFS on graph, starting at node %d\n", start)
	visited := make(map[int]bool)
	path := new(list.LinkedList)
	g.doDFS(start, visited, path)
	return path
}

func (g *Graph) doDFS(vertex int, visited map[int]bool, path *list.LinkedList) {
	// visit vertex
	path.Add(vertex)
	visited[vertex] = true
	fmt.Printf("Visited Vertex: %d\n", vertex)

	// recursively traverse adjacent nodes
	for next := g.AdjNodes[vertex].Head; next != nil; next = next.Next {
		if !visited[next.Value] {
			g.doDFS(next.Value, visited, path)
		}
	}
}
