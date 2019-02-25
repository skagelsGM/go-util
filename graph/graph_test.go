package graph

import (
	"fmt"
	"testing"

	"github.com/skagelsGM/util/list"
)

const ()

// Tests
func TestCreateGraph(t *testing.T) {
	vertices := 5
	g := CreateGraph(vertices)

	if g == nil {
		t.Error("Graph was not created!")
	}

	if g.Vertices != vertices {
		t.Errorf("Graph has incorrect number of vertices: %d, Expected: %d !!!", g.Vertices, vertices)
	}
}

func TestAddEdge(t *testing.T) {
	vertices := 5
	g := CreateGraph(vertices)
	g.AddEdge(0, 1)
	assertAdjacentNodes(t, g, 0, 1)
	assertAdjacentNodes(t, g, 1, 0)
}

func assertAdjacentNodes(t *testing.T, g *Graph, src, dest int) {
	var node = g.AdjNodes[src].Head
	if node == nil {
		t.Errorf("Graph Node %d - Adjacent Nodes Head not set", src)
	}

	if node.Value != dest {
		t.Errorf("Graph Node %d - Adjacent Node %d not found", src, dest)
	}
}

func TestIsAdjacent(t *testing.T) {
	vertices := 5
	g := CreateGraph(vertices)
	edges := [][]int{
		{0, 1},
		{0, 4},
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{3, 4},
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
		assertAdjacent(t, g, edge[0], edge[1])
	}
}

func assertAdjacent(t *testing.T, g *Graph, src, dest int) {
	fmt.Printf("Asserting Adjacent: %d, %d ...\n", src, dest)
	assertTrue(t, g.IsAdjacent(src, dest), fmt.Sprintf("IsAdjacent(%d, %d)", src, dest))
	assertTrue(t, g.IsAdjacent(dest, src), fmt.Sprintf("IsAdjacent(%d, %d)", dest, src))
}

func assertTrue(t *testing.T, check bool, errMsg string) {
	if !check {
		t.Errorf("assertTrue Failed: Expected true; %s", errMsg)
	}
}

func TestBFS(t *testing.T) {
	// 5 nodes, 4 edges: 0-1, 0-2, 0-3, 2-4
	expected := []int{0, 1, 2, 3, 4}
	g := CreateGraph(5)
	edges := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{2, 4},
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	path := g.BreadthFirstSearch(0)
	assertLinkedList(t, path, expected)
}

func TestDFS(t *testing.T) {
	// 5 nodes, 6 edges
	//0-1, 0-2, 0-3, 2-4
	expected := []int{2, 0, 1, 3}
	g := CreateGraph(4)
	edges := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{2, 0},
		{2, 3},
		{3, 3},
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	path := g.DepthFirstSearch(2)
	assertLinkedList(t, path, expected)
}

func assertLinkedList(t *testing.T, list *list.LinkedList, expected []int) {
	next := list.Head
	for _, e := range expected {
		if next.Value != e {
			t.Errorf("Incorrect path: visited node %d, expected: %d", next.Value, e)
		}
		next = next.Next
	}
}
