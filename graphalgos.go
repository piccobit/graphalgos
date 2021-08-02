package graphalgos

import (
	"fmt"
	"strings"
)

// Graph is the basic structure of the graph.
type Graph struct {
	adjacency map[string][]string
}

// NewGraph returns a new Graph.
func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}

// AddVertex adds a vertex to the graph.
func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Printf("vertex %v already exists\n", vertex)
		return false
	}
	g.adjacency[vertex] = []string{}
	return true
}

// AddEdge adds an edge to the vertex.
func (g *Graph) AddEdge(vertex string, node string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Printf("vertex %v does not exists\n", vertex)
		return false
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		fmt.Printf("node %v already exists\n", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return true
}

// BFS is doing a Breadth-First-Search on the graph.
func (g Graph) BFS(startingNode string) []string {
	visited := g.createVisited()
	var q []string
	var r []string

	visited[startingNode] = true
	q = append(q, startingNode)

	for len(q) > 0 {
		var current string
		current, q = q[0], q[1:]
		r = append(r, current)
		for _, node := range g.adjacency[current] {
			if !visited[node] {
				q = append(q, node)
				visited[node] = true
			}
		}
	}

	return r
}

// DFSIterative is doing a Depth-First-Search (iterative) on the graph.
func (g Graph) DFSIterative(startingNode string) []string {
	visited := g.createVisited()
	var s []string
	var r []string

	visited[startingNode] = true
	s = append(s, startingNode)

	for len(s) > 0 {
		var current string
		current, s = s[len(s)-1], s[:len(s)-1]
		r = append(r, current)
		for _, node := range g.adjacency[current] {
			if !visited[node] {
				s = append(s, node)
				visited[node] = true
			}
		}
	}

	return r
}

// DFSRecursive is doing a Depth-First-Search (recursive) on the graph.
func (g Graph) DFSRecursive(startingNode string) []string {
	visited := g.createVisited()
	var result []string

	g.dfsRecursive(startingNode, visited, &result)

	return result
}

func (g Graph) dfsRecursive(startingNode string, visited map[string]bool, nodes *[]string) {
	visited[startingNode] = true
	r := append(*nodes, startingNode)

	for _, node := range g.adjacency[startingNode] {
		if !visited[node] {
			g.dfsRecursive(node, visited, &r)
		}
	}
}

func (g Graph) CreatePath(firstNode, secondNode string) bool {
	visited := g.createVisited()
	var (
		path []string
		q    []string
	)
	q = append(q, firstNode)
	visited[firstNode] = true

	for len(q) > 0 {
		var currentNode string
		currentNode, q = q[0], q[1:]
		path = append(path, currentNode)
		edges := g.adjacency[currentNode]
		if contains(edges, secondNode) {
			path = append(path, secondNode)
			fmt.Println(strings.Join(path, " -> "))
			return true
		}

		for _, node := range g.adjacency[currentNode] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}
	fmt.Println("no link found")
	return false
}

func (g Graph) createVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}
	return visited
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
