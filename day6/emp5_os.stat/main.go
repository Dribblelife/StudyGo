package main

import (
	"os"
	"fmt"
)

func main() {
	list:=os.Args
	if len(list)!=2 {
		fmt.Println("usage:xxx file")
		return
	}
	filename:=list[1]
	info,err:=os.Stat(filename)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("name=",info.Name())
	fmt.Println("size=",info.Size())
}
