package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
}

func main()  {
	var p1 Student
	p1.id=15
	p1.name="zhang"
	p1.sex='m'
	fmt.Println(p1)

	//1.定义一个普通结构体变量
	var s Student
	//定义一个指针变量，保存s地址
	var p2 *Student=&s
	p2.sex='f'
	p2.id=18
	(*p2).name="ma"
	fmt.Println(*p2)
	//2.通过new 申请一个结构体
	p3:=new(Student)
	p3.name="hello"
	p3.id=22
	fmt.Println(*p3)
}
