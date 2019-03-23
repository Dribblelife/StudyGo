package main

import (
	"html/template"
	"fmt"
	"net/http"
	"io"
)

var myTemplate *template.Template

type Result struct {
	output string
}

func (p *Result) Write(b []byte) (n int,err error)  {
	fmt.Println("called by template")
	p.output+=string(b)
	return len(b),nil
}

type Person struct {
	Name string
	Age int
	Title string
}

func userInfo(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("handle hello")
	var arr []Person
	p:=Person{"mac",12,"个人网站"}
	p1:=Person{"mac1",24,"个人网站"}
	p2:=Person{"mac2",18,"个人网站"}
	arr=append(arr,p)
	arr=append(arr,p1)
	arr=append(arr,p2)

	resultWrite:=&Result{}
	io.WriteString(resultWrite,"hello")
	err:=myTemplate.Execute(w,arr)	//把arr中的数据和w绑定
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("template render data:",resultWrite.output)
}

func initTemplate(filename string) (err error)  {
	myTemplate,err=template.ParseFiles(filename)
	if err!=nil {
		fmt.Println(err)
		return
	}
	return
}

func main() {
	initTemplate("F:/GoWorkSpace/src/LearGoProject/day7/emp_newProject/index.html")
	http.HandleFunc("/user/info",userInfo)
	err:=http.ListenAndServe("localhost:8080",nil)
	if err!=nil {
		fmt.Println("http listen failed")

	}

}
