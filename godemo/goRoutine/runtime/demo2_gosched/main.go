package main

import (
	"fmt"
	"runtime"
)

func main()  {
	go func(s string) {
		for i:=0;i<5 ;i++  {
			fmt.Println(s)
		}


	}("world")

	for i:=0;i<5 ;i++  {
		runtime.Gosched()

		fmt.Println("hello")
	}

}
