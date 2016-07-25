package graph

import (
)

type Graph interface {
	AddVertex(name string)
	AddEdge(src string, dst string) error
	SetEdgeProperty(src string, dst string, name string, val int)  error
	GetEdgeProperty(src string, dst string, name string) (int , error)
	GetInEdges(v string) ([]string, error)
	GetOutEdges(v string) ([]string, error)
	GetNumVertices() uint32
	GetNumEdges() uint32
}

