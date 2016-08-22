package graph

type Graph interface {
	AddVertex(name string)
	AddEdge(src string, dst string) error
	SetEdgeProperty(src string, dst string, name string, val interface{}) error
	GetEdgeProperty(src string, dst string, name string) (interface{}, error)
	GetInEdges(v string) ([]string, error)
	GetOutEdges(v string) ([]string, error)
	GetNumVertices() uint32
	GetNumEdges() uint32
}
