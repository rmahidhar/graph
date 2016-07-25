package main

import (
	"fmt"
    graph "github.com/rmahidhar/graph/adjlist"
)

func main() {
	g := graph.New()
	g.AddVertex("SFO")
	g.AddVertex("LA")
	g.AddVertex("NY")
	g.AddVertex("FL")
	g.AddEdge("SFO", "LA")
	g.SetEdgeProperty("SFO", "LA", "BW", 100)
	fmt.Printf("num of vertices %d\n", g.GetNumVertices())
	val, status := g.GetEdgeProperty("SFO", "LA", "BW")
	if status == nil {
		fmt.Printf("SFO -> LA BW %d\n", val)
	}
}
