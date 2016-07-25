package adjlist

import (
	"fmt"
	graph "github.com/rmahidhar/graph"
)

type vertex struct {
    id uint32
	name string
	outEdges map[string]*edge
	inEdges map[string]*edge
	//outEdges []*edge
	//inEdges []*edge
}

type edge struct {
	src string
	dst string
	property map[string]int
}

type AdjList struct {
	numVertices uint32
	numEdges    uint32
	vertexmap   map[string]*vertex
}

func New() graph.Graph {
	return new(AdjList).Init()
}

func (g *AdjList) Init() graph.Graph {
	g.numVertices = 0
	g.numEdges = 0
	g.vertexmap = make(map[string]*vertex)
	return g
}

func (g *AdjList) AddVertex(name string)  {
	g.numVertices += 1
    v := &vertex{
		id: g.numVertices,
		name: name,
		outEdges: make(map[string]*edge),
		inEdges: make(map[string]*edge),
	}
	g.vertexmap[name] = v
}

func (g *AdjList) AddEdge(src string, dst string) error {
	v, srcexist := g.vertexmap[src]
	_, dstexist := g.vertexmap[dst]
	if srcexist && dstexist {
		e := &edge{
			src: src,
			dst: dst,
			property: make(map[string]int),
		}
		//v.outEdges = append(v.outEdges, e)
		//v.inEdges = append(v.inEdges, e)
		v.outEdges[dst] = e
		v.inEdges[dst] = e
		return nil
	} else {
		if (srcexist) {
			return fmt.Errorf("dst vertex %s doesnt exist", dst)
		} else if (dstexist) {
			return fmt.Errorf("src vertex %s doesnt exist", src)
		} else {
			return fmt.Errorf("src %s and dst %s vertex doesnt exist", src, dst)
		}
	}
}

func (g *AdjList) SetEdgeProperty(src string, dst string, name string, val int) error {
	v, srcexist := g.vertexmap[src]
	_, dstexist := g.vertexmap[dst]
	if srcexist && dstexist {
		v.outEdges[dst].property[name] = val
		v.inEdges[dst].property[name] = val
		return nil
	} else {
		if (srcexist) {
			return fmt.Errorf("dst vertex %s doesnt exist", dst)
		} else if (dstexist) {
			return fmt.Errorf("src vertex %s doesnt exist", src)
		} else {
			return fmt.Errorf("src %s and dst %s vertex doesnt exist", src, dst)
		}
	}
}

func (g *AdjList) GetEdgeProperty(src string, dst string, name string) (int, error) {
	v, vexist := g.vertexmap[src]
	if vexist  {
		e, exist := v.outEdges[dst]
		if exist {
			p, valid := e.property[name]
			if valid {
				return p, nil
			} else {
				return -1, fmt.Errorf("property %s doesnt exist", name)
			}
		} else {
			return -1, fmt.Errorf("edge %s -> %s doesnt exist", src, dst)
		}
	} else {
		return -1, fmt.Errorf("src vertex %s doesnt exist", src)
	}
}

func Keys(m map[string]*edge) []string {
/*
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return nil
	}
	return v.MapKeys()
*/
    var keys []string
    for k := range m {
		keys = append(keys, k)
	}
    return keys
}

func (g *AdjList) GetInEdges(name string) ([]string, error) {
	v, exist := g.vertexmap[name]
	if exist {
		return Keys(v.inEdges), nil
	}
	return nil, fmt.Errorf("vertext %s doesn't exist", name)
}

func (g *AdjList) GetOutEdges(name string) ([]string, error) {
	v, exist := g.vertexmap[name]
	if exist {
		return Keys(v.outEdges), nil
	}
	return nil, fmt.Errorf("vertext %s doesn't exist", name)
}

func (g *AdjList) GetNumVertices() uint32 {
	return g.numVertices
}

func (g *AdjList) GetNumEdges() uint32 {
	return g.numEdges
}
