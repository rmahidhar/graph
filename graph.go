package graph

import (
	"container/list"
	//"fmt"
)

type Vertex interface {
	Id() uint32
	Name() string
}

type Property interface {
	Cost() uint32
	Latency() uint32
}

type Edge interface {
	Src() *Vertex
	Dst() *Vertex
	Property() *Property
}

type Graph interface {
	AddVertex(v uint32)
	AddEdge(src uint32, dst uint32, attrs *Property)
	GetNeighbors(v uint32) *list.List
	GetNumVertices() uint32
	GetNumEdges() uint32
}
