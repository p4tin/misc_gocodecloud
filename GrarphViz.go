package main

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
)

func main() {

	g := gographviz.NewGraph()
	g.SetName("G")
	g.SetDir(true)
	g.AddNode("G", "Hello", nil)
	g.AddNode("G", "World", nil)
	g.AddEdge("Hello", "World", true, nil)
	s := g.String()
	fmt.Println(s)
}
