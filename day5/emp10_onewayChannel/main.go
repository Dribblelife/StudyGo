package main

import (
	"fmt"
	)

func producter(out chan<- int)  {//此通道只能写，不能读
	defer close(out)
	for i:=0;i<5 ;i++  {
		out<-i*i
	}

}

func consumer(in <-chan int)  { //只读
	for num:=range in {
		fmt.Println("num=",num)
	}
}

func main() {
	//创建一个双向管道
	ch:=make(chan int)
	//生产者，生出数字，写入管道
	go producter(ch)
	//消费者，从channel读取打印内容
	consumer(ch)
}
