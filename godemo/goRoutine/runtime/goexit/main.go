package main

import (
	"fmt"
	"runtime"
)

func main()  {
	go func() {
		defer fmt.Println("defer A")

		func(){
			defer fmt.Println("defer B")
			runtime.Goexit()
			fmt.Println("B")
		}()

		runtime.Goexit()
		fmt.Println("A")

	}()
	//创建一个死循环，目的不让主goroutine结束
	for true{

	}
}
