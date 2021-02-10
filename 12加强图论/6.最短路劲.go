package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main1() {
	g := graph.New(6)
	g.AddBothCost(0, 1, 8) //  0==1--2
	g.AddBothCost(0, 3, 2) //  |  |  |
	g.AddBothCost(1, 2, 2) //  3--4==5
	g.AddBothCost(1, 4, 2) //
	g.AddBothCost(2, 5, 2) //  -- cost 2
	g.AddBothCost(3, 4, 2) //  == cost 8
	g.AddBothCost(4, 5, 8)

	path, dist := graph.ShortestPath(g, 0, 5)
	fmt.Println("path:", path, "length:", dist)
}

func main() {
	g0 := graph.New(0)
	fmt.Println(g0)

	g1 := graph.New(1)
	g1.Add(0, 0)
	fmt.Println(g1)

	g4 := graph.New(4) //             8
	g4.AddBoth(0, 1)   //  0 <--> 1 <--- 2      3
	g4.AddCost(2, 1, 8)
	fmt.Println(g4)
}