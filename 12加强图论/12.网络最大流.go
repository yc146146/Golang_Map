package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main() {
	g := graph.New(6)
	for _, e := range []struct {
		v, w int
		c    int64
	}{
		{0, 1, 16}, {0, 2, 13}, {1, 2, 10}, {2, 1, 4},
		{1, 3, 12}, {2, 4, 14}, {3, 2, 9}, {4, 3, 7},
		{3, 5, 20}, {4, 5, 4},
	} {
		g.AddCost(e.v, e.w, e.c)
	}

	fmt.Println(graph.MaxFlow(g,0,5))

}
