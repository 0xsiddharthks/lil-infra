package stdlib

import "golang.org/x/exp/slices"

type Directedness rune

const (
	DIRECTED   Directedness = iota
	UNDIRECTED Directedness = iota
)

type edgeList []struct {
	start  int
	end    int
	weight int
}

type adjacencyList [][]struct {
	vertex int
	weight int
}

type Graph struct {
	Directedness Directedness

	EdgeList edgeList

	AdjacencyList adjacencyList

	Matrix [][]int
}

func CreateGraph(edges edgeList, v int) adjacencyList {
	// n(v) - 1 <= n(e) < C(v, 2)
	adjList := make([][]struct {
		vertex int
		weight int
	}, v)

	for i := range adjList {
		adjList[i] = make([]struct {
			vertex int
			weight int
		}, 0)
	}

	for _, e := range edges {
		adjList[e.start] = append(adjList[e.start], struct {
			vertex int
			weight int
		}{e.end, e.weight})
	}

	return adjList

}

func CheckIfEdgeExists(vertices adjacencyList, start int, end int) bool {
	return slices.IndexFunc(vertices[start], func(foo struct {
		vertex int
		weight int
	}) bool {
		return foo.vertex == end
	}) != -1
}

func BFS(vertices adjacencyList) {

}

func DFS(vertices adjacencyList) {}

func CheckIfLoopExists(vertices adjacencyList) bool {
	return false
}
