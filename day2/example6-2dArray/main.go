package main

import "fmt"

func main()  {
	//有多少个[]就是多少维
	//有多少个[]就用多少个循环
	var a[3][4]int
	k:=0

	for i:=0;i<3 ;i++  {
		for j:=0;j<4;j++ {
			k++
			a[i][j]=k
			fmt.Printf("a[%d][%d]=%d ",i,j,a[i][j])
		}
		fmt.Println()
	}

	//二维数组初始化
	b:=[3][4]int{{1,1,2,4},{2:13,3:23},{2}}
	fmt.Println(b)
}
