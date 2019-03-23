package main

import "fmt"

func main()  {
	s:=make([]int,0,3)			//切片容量为3
	oldCap := cap(s)
	for i:=0;i<30;i++  {
		s=append(s,i)
		if newCap:=cap(s);oldCap<newCap {
			fmt.Printf("cap:%d=========>%d\n",oldCap,newCap)
			oldCap=newCap
		}

	}
	//copy使用
	srcSlice:=[]int{2,3,4,5,6}
	dstSlice:=[]int{5,5,5}
	copy(dstSlice,srcSlice)
	fmt.Println(dstSlice)
	fmt.Println("姓名：自己的姓名")
	fmt.Println("性别：自己性别")
	fmt.Println("手机号码：自己的手机号码")
}
