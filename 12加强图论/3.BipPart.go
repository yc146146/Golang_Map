package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main() {
	gm := graph.New(6)
	gm.AddBoth(0, 1) //  0--1--2
	gm.AddBoth(0, 3) //  |  |  |
	gm.AddBoth(1, 2) //  3--4  5
	gm.AddBoth(1, 4)
	gm.AddBoth(2, 5)
	gm.AddBoth(3, 4)
	g := graph.Sort(gm)
	part, ok := graph.Bipartition(g)

	if !ok{
		fmt.Println("fails")
	}

	fmt.Println(part)
}
