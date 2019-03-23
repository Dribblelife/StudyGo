package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i:=0;i<5 ;i++  {
			fmt.Println("go")
		}

	}()

	for i:=0;i<2 ;i++  {
		runtime.Gosched()//让出时间片，先让别的协议执行
		fmt.Println("hello")
	}
}
