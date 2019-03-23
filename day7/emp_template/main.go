package main

import (
	"html/template"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Title string
	age string

}

func main() {
	t,err:=template.ParseFiles("F:/GoWorkSpace/src/LearGoProject/day7/emp_template/index.html")
	if err!=nil{
		fmt.Println(err)
		return
	}
	p:=Person{"mac","个人网站","23"}
	if err:=t.Execute(os.Stdout,p);err!=nil{	//(os.Stdout,p)把p中数据与标准输出绑定
		fmt.Println(err)
	}
}
