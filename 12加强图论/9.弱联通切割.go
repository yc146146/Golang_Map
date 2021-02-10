package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main() {
	g := graph.New(4)


	g.Add(0, 1)
	g.Add(2, 1)
	g.AddBoth(0, 1)
	g.AddBoth(1, 2)
	g.AddBoth(2, 3)
	g.AddBoth(0, 3)
	fmt.Println(graph.Components(g))
}
