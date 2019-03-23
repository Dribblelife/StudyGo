package main

import (
	"net/http"
	"fmt"
)

func handleHello(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("hello handle")
	fmt.Fprint(w,"hello")
}
func login(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("hello,login")
	fmt.Fprint(w,"登录成功了棒槌")
}
func history(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("hello , history" +
		"just do it")
	fmt.Fprint(w,"棒槌，这是记录")
}
func main() {
	http.HandleFunc("/",handleHello)
	http.HandleFunc("/usr/login",login)
	http.HandleFunc("/usr/his",history)
	err:=http.ListenAndServe("localhost:8000",nil)
	if err != nil{
		fmt.Println("http listen failed")
		return
	}
}
