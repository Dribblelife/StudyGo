package main

import (
	"fmt"
	)

/*func exchangeNum(a int,b int)  {				并未改变，是因为两个a、b不一样，不是同一地址的a，b，所以要用指针
	temp:=a
	a=b
	b=temp
	return
}*/

func exchangeNum(a *int,b *int )  {
	var temp int
	temp=*a
	*a=*b
	*b=temp
	return
}

func main()  {
	a:=4
	b:=10
	fmt.Println(a)
	fmt.Println(b)
	exchangeNum(&a,&b)
	fmt.Println(a)
	fmt.Println(b)
}
