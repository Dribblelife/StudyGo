package main

import (
	"fmt"
	"time"
)

func Printer(str string)  {
	for _,data:=range str {
		fmt.Printf("%c",data)
		time.Sleep(500*time.Millisecond)
	}
	fmt.Println()

}

var ch=make(chan int)


func person1(){
	Printer("hello")
	ch<-888
}

func person2(){
	<-ch
	Printer("go")
}

func main() {

	go person1()
	go person2()
	for   {

	}
}
