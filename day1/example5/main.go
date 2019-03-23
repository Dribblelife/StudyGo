package main

import (
	"fmt"
	)

func main()  {
	const LENGTH int  = 5
	const WIDTH  = 10
	var area=LENGTH*WIDTH
	fmt.Printf("面积为：%d",area)
	fmt.Println()
	const a,b,c  = 1,"hh",true
	fmt.Println(a,b,c)
	const(
		unknown=iota		//特殊常量 iota 默认为0 在下一个const出现前 下一个iota自增1 下一行可省略
		female
		male
		)
	fmt.Println(unknown,female,male)
}
