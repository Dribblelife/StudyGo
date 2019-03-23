package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main()  {
	//如果种子参数一致，每次运行程序产生结果一样
	//需要让每次运行种子参数不一样
	//rand.seed(666)
	rand.Seed(time.Now().UnixNano())			//以当前时间作为种子参数
	for i:=0;i<5 ;i++  {
		//fmt.Println("rand=",rand.Int())
		fmt.Println("rand=",rand.Intn(10000))
	}
}
