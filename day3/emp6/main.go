package main

import (
	"fmt"
	)

type Student struct {
	id int
	age int
	name string
}

func main()  {
	s1:=Student{}
	s1.id=1
	s1.name="mac"
	fmt.Println("传参前:",s1)
	test1(s1)
	fmt.Println("传参后",s1)
	test2(&s1)
	fmt.Println("传参后",s1)
}

func test1(s Student)  {	//值传递 不能改变
	s.id=2
	fmt.Println("test1:",s)
}
func test2(s *Student)  {	//地址传递 （引用传递） 形参可以改变实参
	s.id=666
}