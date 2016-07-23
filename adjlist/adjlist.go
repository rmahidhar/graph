package adjlist

import (
	"container/list"
	//"fmt"
	graph "github.com/rmahidhar/graph"
)

type vertex struct {
    id uint32
	name string
	outEdges map[string]edge
	inEdges map[string]edge
}

type edge struct {
	src string 
	dst string
	property map[string]int
}

type AdjList struct {
	numVertices uint32
	numEdges    uint32
	vertexmap   map[string]vertex
}

type VertexMap map[string]vertex

func New() graph.Graph {
	return new(AdjList).Init()
}

func (g *AdjList) Init() graph.Graph {
	g.numVertices = 0
	g.numEdges = 0
	v.vertextmap = make(map[string]vertex)
	return g
}

func (g *AdjList) AddVertex(name string) graph.Vertex {
    v := Vertex{id: ++g.numVertices, name: name}


}

func (g *AdjList) AddEdge(src string, dst string) {

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
