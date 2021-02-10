package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main() {
	g := graph.New(6)
	g.AddBoth(0, 1) //  0 <--> 1 <--> 2
	g.AddBoth(1, 2) //  ^      ^      ^
	g.Add(3, 0)     //  |      |      |
	g.AddBoth(3, 4) //  3 <--> 4 ---> 5
	g.Add(4, 1)
	g.Add(4, 5)
	g.Add(5, 2)

	fmt.Println(graph.StrongComponents(g))
}
