package main

import (
	"LearGoProject/day3/test"
	"fmt"
)

func main()  {
	test.MyFunc()
	s:=test.Stu{}
	s.Id=666
	fmt.Println(s)
}
