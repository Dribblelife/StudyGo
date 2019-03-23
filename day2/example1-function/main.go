package main

import "fmt"

func add(a,b int)int  {
	return a+b
}

/*func main()  {			直接调用main
	d:= add(4,5)
	fmt.Println(d)
}*/

func main()  {
	c:= add					//c打印出为一个地址
	fmt.Println(c)

	sum:=c(4,5)
	fmt.Println(sum)
}