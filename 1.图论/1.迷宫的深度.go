package main

import "fmt"

//坐标
type point struct {
	i int
	j int

}

//方向
var dirs = [4]point{{-1,0},{0,-1},{1,0},{0,1}}

var data = [][]int{
	{0,1,1,0,1,1},
{0,0,0,0,1,1},
{1,1,1,0,1,1},
{0,0,1,0,0,0},
{0,0,1,1,1,0},
{0,0,1,1,1,0},
}


func (p point)add(dir point)point{
	return point{p.i + dir.i,p.j+dir.j}

}

//判断节点是否可以运行
func (p point)at(m[][]int)(int, bool){
	if p.i<0 || p.i>=len(m){
		return 0,false
	}
	if p.j<0 || p.j>=len(m[0]){
		return 0,false
	}
	return m[p.i][p.j], true
}

func walkMaze(maze [][]int, start , end point)[][]int{
	//构造矩阵
	steps := make([][]int, len(maze))
	for i:=range steps{
		//对儿数组每个维度展开一下
		steps[i] = make([]int, len(maze[0]))
	}

	//开始点放入队列
	Stack := []point{start}

	for len(Stack)>0{
		//取得第一个
		cur := Stack[len(Stack)-1]
		//删除第一个
		//Stack = Stack[1:]
		Stack = Stack[:len(Stack)-1]
		for _,dir := range dirs{
			//循环每一个方向
			next := cur.add(dir)
			//判断是否可以走
			nextvalue, ok := next.at(maze)
			if !ok || nextvalue==1{
				continue
			}

			//避免回头路
			nextvalue, ok = next.at(steps)
			if !ok || nextvalue!=0{
				continue
			}
			//避免环路
			if next == start{
				continue
			}
			//可以走的路
			Stack = append(Stack, next)
			mazeValue, ok := cur.at(steps)
			if ok{
				//记录位置
				steps[next.i][next.j] = mazeValue+1
			}
			if end.i == next.i && end.j==next.j{
				fmt.Println("走出来了")
			}
		}
	}

	return steps
}

func main() {


	//fmt.Println(data)
	//fmt.Println(dirs)

	for _, line := range data{
		fmt.Printf("%d\n", line)
	}

	fmt.Println()
	steps := walkMaze(data, point{0,0}, point{len(data)-1,len(data[0])-1})

	for _, line := range steps{
		fmt.Printf("%3d\n", line)
	}


}