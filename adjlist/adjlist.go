package adjlist

import (
	"container/list"
	//"fmt"
	graph "github.com/rmahidhar/graph"
)

type AdjList struct {
	numVertices uint32
	numEdges    uint32
	adjList     []*list.List
}

func New() graph.Graph {
	return new(AdjList).Init()
}

func (g *AdjList) Init() graph.Graph {
	g.numVertices = 0
	g.numEdges = 0
	g.adjList = nil
	return g
}

func (g *AdjList) AddVertex(v uint32) {

}

func (g *AdjList) AddEdge(src uint32, dst uint32, attrs *graph.Property) {

}

func (g *AdjList) GetNeighbors(v uint32) *list.List {
	return nil
}

func (g *AdjList) GetNumVertices() uint32 {
	return g.numVertices
}

func (g *AdjList) GetNumEdges() uint32 {
	return g.numEdges
}
