package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main1() {
	g := graph.New(4)
	g.AddBoth(0, 1)
	g.Add(1, 2)

	fmt.Println(graph.EulerDirected(g))

	g1 := graph.New(4)
	g1.Add(0, 1)
	g1.Add(1, 2)
	g1.Add(0, 3)

	fmt.Println(graph.EulerDirected(g1))
}

func main() {
	g := graph.New(4)
	g.AddBoth(1, 3)
	g.AddBoth(2, 3)
	g.AddBoth(3, 3)

	fmt.Println(graph.EulerUndirected(g))
}