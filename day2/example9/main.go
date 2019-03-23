package main

import "fmt"

func main()  {
	a:=[6] int{1,2,3,4,5,6}
	//modify(a)
	modify(&a)
	fmt.Println("main a is:",a)
}

func modify(p *[6] int)  {
	(*p)[0]=666666666
	fmt.Println("modify a is :",*p)
}