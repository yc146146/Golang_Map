package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main() {
	g := graph.New(5)
	g.Add(0, 1)
	g.Add(2, 3)
	g5  := graph.Sort(g)

	g.AddCost(2, 3, 1)
	g5c := graph.Sort(g)

	res := graph.Transpose(g5)
	fmt.Println(res, g5c)
}
