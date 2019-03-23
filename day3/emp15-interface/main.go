package main

import "fmt"

type humaner interface {
	sayHi()
}

type stu struct {
	id int
	name string
}

func (s *stu) sayHi() {
	fmt.Printf("id:%d&name:%s\n",s.id,s.name)
}

type teacher struct {
	addr int
	group int
}

func (s *teacher) sayHi() {
	fmt.Printf("id:%d&name:%d\n",s.addr,s.group)
}

type myStr string

func (s *myStr) sayHi()  {
	fmt.Printf("myStr:%s sayhi\n",*s)
}

//定义一个函数，函数参数为接口类型
//多态：只有一个函数，可以有不同表现
func WhoSayHi(p humaner)  {
	p.sayHi()
}

func main() {
	var p  humaner
	s:=&stu{18,"mark"}
	p=s
	p.sayHi()

	s1:=&teacher{18,19}
	p=s1
	p.sayHi()


	var str myStr ="hello world"

	str.sayHi()

	WhoSayHi(s)
	WhoSayHi(s1)
	WhoSayHi(&str)

	x:=make([]humaner,3)
	x[0]=s
	x[1]=s1
	x[2]=&str
//切片中第一个值返回切片下标，第二个返回下标对应的值
	for _,i:=range x{
		i.sayHi()
	}
}
