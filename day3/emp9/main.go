package main

import "fmt"

type Person struct {
	id int
	name string
}
type Stu struct {
	*Person
	score int
	class string
}

func main()  {
	//顺序初始化
	var s1=Stu{&Person{1,"mac"},100,"math"}
	fmt.Println(s1)
	//自动推导类型
	s2:=Stu{&Person{1,"mac"},100,"math"}
	fmt.Printf("%+v\n",s2)	//%+v 更详细的显示
	//部分初始化
	s3:= Stu{Person:&Person{id:2},class:"Eng"}
	fmt.Printf("%+v\n",s3)

	s4:=new(Stu)
	s4.Person.id=2
	s4.score=100
	fmt.Printf("%+v\n",s4)

}
