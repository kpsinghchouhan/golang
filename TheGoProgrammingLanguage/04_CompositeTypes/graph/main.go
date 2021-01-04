package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("Bangalore", "New Delhi")
	addEdge("Bangalore", "Mumbai")
	addEdge("Bangalore", "Kolkata")
	addEdge("Bangalore", "Chennai")
	fmt.Println(hasEdge("Bangalore", "Mumbai"))
	fmt.Println(hasEdge("Bangalore", "Hyderabad"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
