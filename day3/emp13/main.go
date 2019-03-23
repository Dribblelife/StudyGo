package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	sex byte
}

//接受者为普通变量，非指针,值语义，只是一份拷贝
func (tmp Person) setInfoValue(n string,a int,s byte ){
	tmp.name=n
	tmp.sex=s
	tmp.age=a
	fmt.Println("setInfoValue:",tmp)
}

//接受者作为指针变量，引用传递
func (tmp *Person) setInfoPointer(n string,a int,s byte )  {
	tmp.name=n
	tmp.sex=s
	tmp.age=a
}

func main() {
	var p Person=Person{"mike",1,'f'}
	fmt.Println(p)
	p1:=Person{"mac",18,'f'}
	p1.setInfoValue("go",23,'f')
	fmt.Println(p1)
	p1.setInfoPointer("golang",11,'h')
	fmt.Println(p1)
	}
