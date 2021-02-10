package main

import (
	"fmt"
	"math"
)

var Nodes = 6
var LineCounts = 9
var MyMaps = [9][3]int{
	{2, 4, 11},
	{3, 5, 13},
	{4, 6, 3},
	{5, 6, 4},
	{2, 3, 6},
	{4, 5, 7},
	{1, 2, 1},
	{3, 4, 9},
	{1, 3, 2},
}

func main() {
	//标记路径
	var book [7]int
	//标记距离
	var dis [7]int
	var myMap [7][7]int
	//极大值
	inf := math.MaxInt64
	var count, sum int
	for i := 0; i <= Nodes; i++ {
		for j := 0; j <= Nodes; j++ {
			if i == j {
				myMap[i][j] = 0
			} else {
				myMap[i][j] = inf
			}
		}
	}
	fmt.Println(myMap)
	for _, v := range MyMaps {
		myMap[v[0]][v[1]] = v[2]
		myMap[v[1]][v[0]] = v[2]
	}
	for _, line := range myMap {
		fmt.Println(line)
	}

	for i := 1; i <= Nodes; i++ {
		dis[i] = myMap[1][i]
	}
	book[1] = 1
	count++
	j := 0
	for count < Nodes {
		min := inf
		for i := 1; i <= Nodes; i++ {
			if book[i] == 0 && dis[i] < min {
				min = dis[i]
				j = i
			}

		}
		book[j] = 1
		count++
		sum += dis[j]
		for i, v := range dis {
			if book[i] == 0 && v > myMap[j][i] {
				dis[i] = myMap[j][i]
			}

		}
	}

	fmt.Println(count, sum)
}
