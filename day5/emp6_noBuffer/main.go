package main

import (
	"fmt"
	"time"
)

func main() {
	//创建管道
	ch:=make(chan string)
	defer fmt.Println("主协程结束")
	go func() {
		defer fmt.Println("子协程结束")
		for i:= 0;i<2;i++{
			fmt.Println("我是子协程，i=",i)
		}
		time.Sleep(time.Second)
		ch<-"我是子协程，要工作结束"
		}()
	str:=<-ch
	fmt.Println("str=",str)

}
