package main

import "fmt"

func main()  {
	a:=[9]int{0,1,2,3,4,5,6,7,8}

	s1:=a[2:5]
	s1[2]=666
	s2:=s1[2:7]
	s2[2]=777

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(a)

	}
