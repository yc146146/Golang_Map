package main

import (
	"container/list"
	"fmt"
)

const Max_Szie = 5
const Max_VALUE = 9

type Graph struct {
	vexs    []string
	vexnum  int
	edgenum int
	matrix  [Max_Szie][Max_Szie]int
}

func initGG(gg *Graph) {
	gg.matrix[0][1] = 5
	gg.matrix[0][2] = 3

	gg.matrix[1][0] = 5
	gg.matrix[1][3] = 7
	gg.matrix[1][4] = 4

	gg.matrix[2][0] = 3
	gg.matrix[2][3] = 6

	gg.matrix[3][1] = 7
	gg.matrix[3][2] = 6
	gg.matrix[3][4] = 1

	gg.matrix[4][1] = 4
	gg.matrix[4][3] = 1
	gg.edgenum = 12 / 2
}

func DFS(gg *Graph, visit *[]bool, i int) {
	fmt.Println(gg.vexs[i])
	for j := 0; j < gg.vexnum; j++ {
		if gg.matrix[i][j] != Max_VALUE && !(*visit)[j]{
			(*visit)[j] = true
			DFS(gg, visit, j)
		}
	}
}

func fDFS(gg *Graph) {
	visit := make([]bool, 10, 10)
	fmt.Println(visit)
	visit[0] = true
	DFS(gg, &visit, 0)
}

func fBFS(gg *Graph) {
	listq := list.New()
	visit := make([]bool, 10, 10)

	visit[0] = true
	listq.PushBack(0)

	for listq.Len() > 0 {
		index := listq.Front()
		fmt.Println(gg.vexs[index.Value.(int)])
		for j := 0; j < gg.vexnum; j++ {
			if !visit[j] && gg.matrix[index.Value.(int)][j] != Max_VALUE {
				visit[j] = true
				listq.PushBack(j)
			}
		}
		listq.Remove(index)
	}
}

//打印图
func printG(gg Graph, l int) {
	for i := 0; i < l; i++ {
		fmt.Println(gg.matrix[i])
	}
}

//抓取节点位置
func get_position(gg *Graph, ch string) int {
	for i := 0; i < gg.vexnum; i++ {
		if gg.vexs[i] == ch {
			return i
		}
	}
	return -1
}

func Prim(gg *Graph, start int) {
	index := 0
	sum := 0

	prims := make([]string, 10, 10)
	var weights [5][2]int

	//	设定开始
	prims[index] = gg.vexs[start]
	index++

	for i := 0; i < gg.vexnum; i++ {
		weights[i][0] = start
		weights[i][1] = gg.matrix[start][i]
	}
	//	删除1个
	weights[start][1] = 0

	//算法核心
	for i := 0; i < gg.vexnum; i++ {
		if start == i {
			//重合
			continue
		}
		min := Max_VALUE
		next := 0
		for j := 0; j < gg.vexnum; j++ {
			if weights[j][1] != 0 && weights[j][1] < min {
				min = weights[j][1]
				next = j
			}
		}
		fmt.Println(gg.vexs[weights[next][0]], gg.vexs[next], "权重", weights[next][1])

		//叠加
		sum += weights[next][1]
		prims[index] = gg.vexs[next]

		index++
		//加入节点了删除
		weights[next][1] = 0
		for j := 0; j < gg.vexnum; j++ {
			if weights[j][1] != 0 && gg.matrix[next][j] < weights[j][1] {
				weights[j][1] = gg.matrix[next][j]
				weights[j][0] = next
			}
		}

	}

	fmt.Println("sum", sum)
	fmt.Println(prims)

}

func main() {
	var gg Graph
	var vexs = []string{"B","A","C","D","E"}
	gg.vexnum=5
	gg.vexs=vexs
	for i := 0; i < len(vexs); i++ {
		for j := 0; j < len(vexs); j++ {
			gg.matrix[i][j]=Max_VALUE
		}
	}
	initGG(&gg)
	fmt.Println(gg.vexs)
	fBFS(&gg)
	fDFS(&gg)
	Prim(&gg,0)
}
