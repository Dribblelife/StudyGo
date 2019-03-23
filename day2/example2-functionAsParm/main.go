package main

import (
	"fmt"
	)

type caculator_func func(int,int) int			//type一般定义结构体和接口

func add(a,b int)int  {
	return a+b
}
func sub(a,b int)int  {
	return a-b
}
func operator(op caculator_func,a int,b int) int  {
	return op(a,b)
}
func main()  {

	//sum:=operator(sub,5,6)		一种调用方法，也可以直接使用匿名函数，如下
	sum:= operator(func(a int, b int) int { 		//匿名函数方法
		return a*b
	},5,6)
	fmt.Println(sum)
}
