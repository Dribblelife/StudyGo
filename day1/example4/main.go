package main

import "fmt"

var (
	aa=22						//全局变量简写方式
	bb="hh"
	cc=true

	ddd,bbb=12,"sss"			//使用并列的方式以及类型推断
	//ccc：=12  错误  全局变量不支持：=这种方式

)

func varDef()  {
	var a int
	var b string
	var c bool
	fmt.Printf("%d %q %v",a,b,c)
	fmt.Println()
}

func varDef1()  {
	var a,b int  = 4,5
	var c string  = "nihao"
	var d bool= true
	fmt.Println(a,b,c,d)
}

func varDef2()  {
	var a,b,c,d=2,3,"hello",false
	fmt.Println(a,b,c,d)
}

func varDef3()  {
	a,b,c,d:=2,3,"hello",false			//简写方法
	c="ddddd"                           //重新赋值 不加：
	fmt.Println(a,b,c,d)
}

func main()  {
	varDef()
	varDef1()
	varDef2()
	varDef3()
}
