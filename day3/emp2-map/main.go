package main

import "fmt"

func main()  {
	var m1 map[int]string

	fmt.Println(m1==nil)
	fmt.Println(len(m1))

	m2:=make(map[int]string,2)
	m2[1]="go"
	m2[2]="c"
	m2[3]="java"
	fmt.Println(m2)
	fmt.Println(len(m2))

	m4:= map[int]string{1:"go",2:"c",3:"java"}
	fmt.Println(m4)
	fmt.Println(len(m4))
}
