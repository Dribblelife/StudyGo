package main

import "fmt"

func testA()  {
	fmt.Println("aaaa")
}

func testB(x int)  {
	//fmt.Println("bbbb")
	//显式调用panic
	//panic("this is a panic test")
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(recover())
		}
	}()
	var a[10]int
	a[x]=80
	fmt.Println(a)
}
func testC()  {
	fmt.Println("ccc")
}

func main() {
	testA()
	testB(9)
	testC()
}
