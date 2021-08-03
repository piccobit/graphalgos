package graphalgos

import (
	"fmt"
	"strings"
)

var debug bool

// Graph is the basic structure of the graph.
type Graph struct {
	adjacency map[string][]string
}

// Debug allows to switch on or of the debugging output.
func Debug(state bool) {
	debug = state
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
		if debug {
			fmt.Printf("vertex %v already exists\n", vertex)
		}

		return false
	}
	g.adjacency[vertex] = []string{}
	return true
}

// AddEdge adds an edge to the vertex.
func (g *Graph) AddEdge(vertex string, node string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		if debug {
			fmt.Printf("vertex %v does not exists\n", vertex)
		}

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

	result = g.dfsRecursive(startingNode, visited, &result)

	return result
}

func (g Graph) dfsRecursive(startingNode string, visited map[string]bool, result *[]string) []string {
	visited[startingNode] = true
	r := append(*result, startingNode)

	for _, node := range g.adjacency[startingNode] {
		if !visited[node] {
			r = g.dfsRecursive(node, visited, &r)
		}
	}

	return r
}

func (g Graph) CreatePath(firstNode, secondNode string, reverse bool) ([]string, bool) {
	visited := g.createVisited()
	var (
		path []string
		q    []string
	)

	if reverse {
		tmpNode := firstNode
		secondNode = firstNode
		firstNode = tmpNode
	}

	if debug {
		if reverse {
			fmt.Println("Direction: reverse")
		} else {
			fmt.Println("Direction: normal")
		}
	}

	q = append(q, firstNode)
	visited[firstNode] = true

	for len(q) > 0 {
		var currentNode string
		currentNode, q = q[0], q[1:]
		path = append(path, currentNode)
		edges := g.adjacency[currentNode]
		if contains(edges, secondNode) {
			path = append(path, secondNode)

			if debug {
				fmt.Println(strings.Join(path, " -> "))
			}

			return path, true
		}

		for _, node := range g.adjacency[currentNode] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}

	if debug {
		fmt.Println("no link found")
	}

	return nil, false
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

func (g Graph) IsLeaf(node string) bool {
	if len(g.adjacency[node]) != 0 {
		return false
	}

	return true
}
