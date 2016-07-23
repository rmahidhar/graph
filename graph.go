package graph

import (
	"container/list"
	//"fmt"
)

type Vertex interface {
	Id() uint32
	Name() string
}

type Edge interface {
	Src() *Vertex
	Dst() *Vertex
	Property(name string) uint32
}

type Graph interface {
	AddVertex(name string)
	AddEdge(src string, dst string)
	SetEdgeProperty(src string, dst string, name string, val uint32)
	GetInEdges(v string) map[string]Edge
	GetOutEdges(v string) map[string]Edge
	GetNumVertices() uint32
	GetNumEdges() uint32
}
