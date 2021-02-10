package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main() {
	g := graph.New(10)
	g.AddBothCost(0, 1, 4)
	g.AddBothCost(0, 7, 8)
	g.AddBothCost(1, 2, 8)
	g.AddBothCost(1, 7, 11)
	g.AddBothCost(2, 3, 7)
	g.AddBothCost(2, 8, 2)
	g.AddBothCost(2, 5, 4)
	g.AddBothCost(3, 4, 9)
	g.AddBothCost(3, 5, 14)
	g.AddBothCost(4, 5, 10)
	g.AddBothCost(5, 6, 2)
	g.AddBothCost(6, 7, 1)
	g.AddBothCost(6, 8, 6)
	g.AddBothCost(7, 8, 7)


	fmt.Println(graph.MST(g))
}
