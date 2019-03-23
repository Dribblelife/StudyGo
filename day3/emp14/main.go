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
func (tmp Person) setInfoValue(){

	fmt.Println("setInfoValue:")
}

//接受者作为指针变量，引用传递
func (tmp *Person) setInfoPointer()  {
	fmt.Println("setInfoPointer")
}

func main() {
	//指针变量，他能调用哪些方法，这些方法的集合称为方法集
	p:=&Person{"mile",1,'g'}
	p.setInfoPointer()
	//内部做了转换
	p.setInfoValue()

}
