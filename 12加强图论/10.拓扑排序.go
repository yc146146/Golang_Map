package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main() {
	g := graph.New(5)
	g.Add(0, 1)
	g.Add(1, 2)
	g.Add(1, 3)
	g.Add(2, 4)
	g.Add(3, 4)
	order, ok := graph.TopSort(g)
	fmt.Println(order,ok)


}
