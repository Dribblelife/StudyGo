package main

import "fmt"

func main()  {
	a:=[5]int{}

	fmt.Printf("len=%d,cap=%d\n",len(a),cap(a))

	s:=[]int{}
	fmt.Printf("len=%d,cap=%d\n",len(s),cap(s))
	s=append(s,2)
	fmt.Printf("len=%d,cap=%d\n",len(s),cap(s))
	//几种切片的创建方式
	//1.自动推导类型，并初始化
	s1:=[]int{1,2,3,4}
	fmt.Println(s1)
	//2.借助make 格式：make（切片类型，长度，容量） 未指定容量时与长度相等
	s2:=make([]int,5,7)
	fmt.Printf("len=%d,cap=%d\n",len(s2),cap(s2))

	s3:=make([]int,5)
	fmt.Printf("len=%d,cap=%d\n",len(s3),cap(s3))
}
