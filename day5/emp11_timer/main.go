package main

import (
		"fmt"
	"time"
	)

func time1()  {
	//只验证一次
	timer:=time.NewTimer(1*time.Second)
	for   {
		<-timer.C
		fmt.Println("时间到")
	}
}

func afterTime(){
	//延时
	<-time.After(2*time.Second)
	fmt.Println("时间到")
}

func main() {
	//创建一个定时器，设置时间为2s，2s后往time通道写内容（当前时间）
	timer:=time.NewTimer(2*time.Second)
	fmt.Println("当前时间为：",time.Now())
	//2s后，往time.C里面写数据，有数据后，就可以读取
	t:=<-timer.C	//channel没有数据前阻塞
	fmt.Println("t=",t)

	//time1()

	afterTime()
}
