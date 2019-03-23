package main

import "fmt"

func add(a,b int)int  {
	return a+b
}

type myInt int

func (a myInt) add1(b myInt)myInt  {
	return a+b
}

func main()  {
	var result1 int
	result1=add(1,3)
	fmt.Println(result1)
	var c myInt =20
	result2:=c.add1(10)
	fmt.Println(result2)
}
