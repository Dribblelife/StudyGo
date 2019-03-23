package main

import (
	"fmt"
	)

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string

}


func main()  {
	s1:=Student{1,"mac",'m',18,"xian"}
	fmt.Println(s1)
	s2:=Student{id:2,addr:"beijing"}
	fmt.Println(s2)

	//指针类型两种初始化
	var p1 *Student  = &Student{0,"mac",'m',18,"xian"}

	p2:=&Student{name:"jaten"}

	fmt.Println(*p1)
	fmt.Printf("p1 type is %T\n",p1)
	fmt.Println(*p2)
	fmt.Printf("p2 type is %T",p2)
}
