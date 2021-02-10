package main

import (
	"fmt"
	"yinchen.com/文件4/12加强图论/graph"
)

func main() {
	// Build a graph with four vertices and four undirected edges.
	// (Each of these edges are, in fact, represented by two directed
	// edges pointing in opposite directions.)
	g := graph.New(4)
	g.AddBoth(0, 1) //  0 -- 1
	g.AddBoth(0, 2) //  |    |
	g.AddBoth(2, 3) //  2 -- 3
	g.AddBoth(1, 3)

	// The vertices of all graphs in this package are numbered 0..n-1.
	// The edge iterator is a method called Visit; it calls a function
	// for each neighbor of a given vertex. Together with the Order
	// method — which returns the number of vertices in a graph — it
	// constitutes an Iterator. All algorithms in this package operate
	// on any graph implementing this interface.

	// Visit all edges of a graph.
	for v := 0; v < g.Order(); v++ {
		g.Visit(v, func(w int, c int64) (skip bool) {
			// Visiting edge (v, w) of cost c.
			return
		})
	}

	// The immutable data structure created by Sort has an Iterator
	// that returns neighbors in increasing order.

	// Visit the edges in order.
	for v := 0; v < g.Order(); v++ {
		graph.Sort(g).Visit(v, func(w int, c int64) (skip bool) {
			// Visiting edge (v, w) of cost c.
			return
		})
	}

	// The return value of an iterator function is used to break
	// out of the iteration. Visit, in turn, returns a boolean
	// indicating if it was aborted.

	// Skip the iteration at the first edge (v, w) with w equal to 3.
	for v := 0; v < g.Order(); v++ {
		aborted := graph.Sort(g).Visit(v, func(w int, c int64) (skip bool) {
			fmt.Println(v, w)
			if w == 3 {
				skip = true // Aborts the call to Visit.
			}
			return
		})
		if aborted {
			break
		}
	}
}