package main

import (
	"runtime"
	"fmt"
)

func main(){
	n:=runtime.GOMAXPROCS(4)
	fmt.Printf("n=%d\n",n)
	for i:=0;i<50 ;i++  {
		go fmt.Println(1)
		fmt.Println(0)
	}
}
