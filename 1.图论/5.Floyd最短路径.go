package main

import "fmt"

var pointNum = 4
var Pashroad = [8][3]int{
	{1,2,2},
	{1,3,6},
	{1,4,4},
	{2,3,3},
	{3,1,7},
	{3,4,1},
	{4,1,5},
	{4,3,12},
}



func main() {
	var mymap[4][4] int
	for i:=0;i<pointNum;i++{
		for j:=0;j<pointNum;j++{
			mymap[i][j]=-1
		}
	}

	fmt.Println("地图尚未初始化")
	for _,line := range mymap{
		fmt.Println(line)
	}
	fmt.Println("地图初始化")

	for _,line := range Pashroad{
		mymap[line[0]-1][line[1]-1]=line[2]
	}



	for _,line := range mymap{
		fmt.Println(line)
	}

	fmt.Println("计算最短路径")

	for k:=0;k<pointNum;k++{
		for i:=0;i<pointNum;i++{
			for j:=0;j<pointNum;j++{
				if i==j{
					continue
				}
				if mymap[i][k] == -1||mymap[k][j] == -1{
					continue
				}
				if mymap[i][k] == -1 || mymap[i][j] > mymap[i][k]+mymap[k][j] {
					mymap[i][j]=mymap[i][k]+mymap[k][j]
				}


			}
		}
	}

	fmt.Printf("%3d\n",[]int{1,2,3,4})
	for i,line := range mymap{
		fmt.Print("",i+1)
		fmt.Printf("%3d\n", line)
	}

}
