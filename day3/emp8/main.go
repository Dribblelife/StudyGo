package main

import "fmt"

type person struct {
	id int
	name string
}
type Stu struct {
	person
	score int
	class string
}

func main()  {
	//顺序初始化
	var s1=Stu{person{1,"mac"},100,"math"}
	fmt.Println(s1)
	//自动推导类型
	s2:=Stu{person{1,"mac"},100,"math"}
	fmt.Printf("%+v\n",s2)	//%+v 更详细的显示
	//部分初始化
	s3:= Stu{score:100}
	fmt.Printf("%+v\n",s3)
}
