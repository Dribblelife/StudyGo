package main

import (
	"fmt"
	"time"
)

func main() {
	//创建无缓冲管道
	ch:=make(chan int,3)
	//内置函数len返回未被读取的缓冲元素，cap返回缓冲区大小
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n",len(ch),cap(ch))
	go func() {
		defer fmt.Println("子协程结束")
		for i:= 0;i<10;i++{
			ch<-i
			fmt.Printf("子协程正在运行[%d]:len(ch)=%d,cap(ch)=%d\n",i,len(ch),cap(ch))
		}
		}()
	time.Sleep(2*time.Second)//延时2s

	for i:=0;i<10 ;i++  {
		num:=<-ch
		fmt.Println("num=",num)
	}

	fmt.Println("main协程结束")

}
