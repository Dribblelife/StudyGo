package main

import (
	"fmt"
	)

func main()  {

	m1:= map[int]string{1:"go",2:"c"}

	fmt.Println(m1)

	m1[1]="python"

	m1[5]="java"
	fmt.Println(m1)

	//map遍历，遍历无序
	for key,value:=range m1{
		fmt.Printf("%d==>%s\n",key,value)
	}
	//查找 map中时候存在某个key
	value,ok:=m1[1]
	if ok==true{
		fmt.Println("该值为：",value)
	}else {
		fmt.Println("key不存在")
	}
	//删除
	delete(m1,1)
	fmt.Println(m1)


	del(m1)
	fmt.Println(m1)


}
	//map作为函数参数，为引用传递
func del(m map[int]string)  {
	delete(m,2)
}