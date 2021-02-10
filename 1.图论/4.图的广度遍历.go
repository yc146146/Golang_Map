package main

import "fmt"

//节点的数量
var pointCount = 5

//变长为8
var LineCount = 7
var Start = 1
var End = 5

//保存最短路径
var Min = -1

var data = [7][3]int{
	{1, 2, 1},
	{1, 5, 1},
	{2, 3, 1},
	{2, 5, 1},
	{3, 1, 1},
	{3, 4, 1},
	{4, 5, 1},
}

type node struct {
	x int
	y int
}

//走过的路
var path = make([]int, 0)

//最短的路
var minpath = make([]int, 0)

//记录边长与路径
var book = make(map[int]int)

//记录所有的可能 记录路劲
var chart = map[int]map[int]int{}

//队列
var Queuelist = make(map[int]node)

func Datainit() {
	for i := 1; i <= pointCount; i++ {
		//每个元素都要初始化
		chart[i] = map[int]int{}
		for j := 1; j <= pointCount; j++ {
			if i == j {
				chart[i][j] = 0
			} else {
				chart[i][j] = -1
			}
		}
	}
	for _, d := range data {
		//两点放入map
		//无向图
		chart[d[0]][d[1]] = d[2]
		chart[d[1]][d[0]] = d[2]
	}
}

func main() {
	Datainit()

	header := 0
	tail := 0
	//插入节点
	Queuelist[tail] = node{Start, 0}

	book[Start] = 1
	//记录路径
	path = append(path, Start)
	tail++

	isEnd := false
	for header < tail {
		//获得当前节点
		nowPoint := Queuelist[header]
		for nextpoint := 1; nextpoint <= pointCount; nextpoint++ {
			if book[nextpoint] == 0 && chart[nowPoint.x][nextpoint] > 0 {
				Queuelist[tail] = node{nextpoint, nowPoint.y + chart[nowPoint.x][nextpoint]}
				tail++
				book[nextpoint] = 1
				path = append(path, nextpoint)
			}
			if nowPoint.x==End{
				isEnd=true
				break
			}
		}
		if isEnd {
			break
		}
		header++
	}

	//最短距离
	fmt.Printf("%d\n", Queuelist[tail-1].y)
	fmt.Println(path)
	fmt.Println(Queuelist)
}
