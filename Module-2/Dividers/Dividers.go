package main

import "fmt"

type Vertex struct {
	v int
}

type Edge struct {
	a, b int
}

type Graph struct {
	vertexes []Vertex
	edges    []Edge
}

func (graph *Graph) AppendVertex(vertex Vertex) {
	graph.vertexes = append(graph.vertexes, vertex)
}

func (graph *Graph) AppendEdge(edge Edge) {
	graph.edges = append(graph.edges, edge)
}

func (graph *Graph) DeleteEdge(edge Edge) {
	i, contain := graph.ContainsEgge(edge)
	if !contain {
		return
	}
	copy(graph.edges[i:], graph.edges[i+1:])
	graph.edges = graph.edges[:len(graph.edges)-1]
}

func (graph *Graph) countVertexes() int {
	return len(graph.vertexes)
}

func (graph *Graph) ContainsEgge(edge Edge) (int, bool) {
	for i := 0; i < len(graph.edges); i++ {
		if edge == graph.edges[i] {
			return i, true
		}
	}
	return -1, false
}

func main() {
	x := myScan()
	graphDivisor := makeGraphDivisor(x)
	printDOT(graphDivisor)
}

func myScan() (res int) {
	_, err := fmt.Scanln(&res)
	if err != nil {
		fmt.Printf("Error: myScan")
		return
	}
	return
}

func printDOT(graph Graph) {
	fmt.Println("graph {")
	for _, vertex := range graph.vertexes {
		fmt.Print("    ")
		fmt.Println(vertex.v)
	}
	for _, edge := range graph.edges {
		fmt.Printf("    %d--%d\n", edge.a, edge.b)
	}
	fmt.Println("}")

}

func makeGraphDivisor(x int) (graphDivisor Graph) {
	setVertexesDivisor(&graphDivisor, x)
	setEdgeDivisor(&graphDivisor)
	return
}

func setVertexesDivisor(graphDivisor *Graph, x int) {
	for i := 1; i <= x/2; i++ {
		if x%i == 0 {
			graphDivisor.AppendVertex(Vertex{x / i})
		}
	}
	graphDivisor.AppendVertex(Vertex{1})
}

func setEdgeDivisor(graphDivisor *Graph) {
	countVertexesGraphDivisor := graphDivisor.countVertexes()
	for i := 0; i < countVertexesGraphDivisor; i++ {
		vertexA := graphDivisor.vertexes[i]
		for j := i + 1; j < countVertexesGraphDivisor; j++ {
			vertexB := graphDivisor.vertexes[j]
			if vertexA.v%vertexB.v == 0 && !isVertexW(graphDivisor.vertexes, vertexA, vertexB) {
				graphDivisor.AppendEdge(Edge{vertexA.v, vertexB.v})
			}
		}
	}
}

func isVertexW(vertexes []Vertex, vertexU Vertex, vertexV Vertex) bool {
	for _, vertexW := range vertexes {
		if vertexW == vertexV || vertexW == vertexU {
			continue
		}
		if vertexU.v%vertexW.v == 0 && vertexW.v%vertexV.v == 0 {
			return true
		}
	}
	return false
}
