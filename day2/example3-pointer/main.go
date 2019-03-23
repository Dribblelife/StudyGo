package main

import "fmt"

func modify(p *int)  {
	fmt.Println(p)
	*p=100
}

func main()  {
	a:=10
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)

	var b int  = 200
	var p *int  = &b
	fmt.Println(*p)

	*p=1000
	fmt.Println(b)

	var c int  = 999
	p = &c
	*p=111
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

