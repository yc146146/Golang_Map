package main

import "fmt"

//节点的数量
var pointCount = 5
//变长为8
var LineCount = 8
var Start =1
var End =5
//保存最短路径
var Min = -1

var data = [8][3]int{
	{1,2,2},
	{1,5,10},
	{2,3,3},
	{2,5,7},
	{3,1,4},
	{3,4,4},
	{4,5,5},
	{5,3,3},
}

//走过的路
var path = make([]int, 0)
//最短的路
var minpath=make([]int,0)
//记录边长与路径
var book=make(map[int]int)
//记录所有的可能 记录路劲
var chart = map[int]map[int]int{}


func Datainit(){
	for i:=1;i<=pointCount;i++{
		//每个元素都要初始化
		chart[i] = map[int]int{}
		for j:=1;j<=pointCount;j++{
			if i==j{
				chart[i][j] = 0
			}else{
				chart[i][j]=-1
			}
		}
	}
	for _,d := range data{
		//两点放入map
		chart[d[0]][d[1]]=d[2]
	}
}

//深度遍历
func DFS(point, nowDis int){
	fmt.Println(minpath,path)

	if Min !=  -1 && nowDis>Min{
		return
	}

	if point == End{
		for Min == -1 || nowDis<Min{
			Min=nowDis
			//记录最短路劲
			minpath = path
			return
		}
		return
	}
	for nextPoint := 1;nextPoint<= pointCount;nextPoint++{
		if book[nextPoint] == 0 && chart[point][nextPoint]>0{
			book[nextPoint]=1
			path = append(path, nextPoint)
			//递归遍历
			DFS(nextPoint, nowDis+chart[point][nextPoint])
			book[nextPoint] = 0
			path=path[:len(path)-1]
		}
	}
}

func main() {
	Datainit()
	book[Start]=1
	//记录路径
	path = append(path, Start)
	DFS(Start, 0)
	//最短距离
	fmt.Printf("%d\n", Min)
	fmt.Println(minpath)
}
