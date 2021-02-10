package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

//bfs.go 广度遍历




func main() {
	gm := graph.New(6)
	gm.AddBoth(0, 1) //  0--1--2
	gm.AddBoth(0, 3) //  |  |  |
	gm.AddBoth(1, 2) //  3--4  5
	gm.AddBoth(1, 4)
	gm.AddBoth(2, 5)
	gm.AddBoth(3, 4)
	g := graph.Sort(gm)

	dist := make([]int, g.Order())
	graph.BFS(g, 0, func(v, w int, _ int64) {
		fmt.Println(v, "to", w)
		dist[w] = dist[v] + 1
	})
	fmt.Println("dist:", dist)
}
