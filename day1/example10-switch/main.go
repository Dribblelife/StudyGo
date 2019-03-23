package main

import "fmt"

func grade(score int)string  {
	g:=""
	switch  {
	case score<0 || score>100:
		panic(fmt.Sprintf("wrong score:%d",score))
	case score<60:
		g = "d"
	case score<80:
		g = "c"
	case score <90:
		g = "b"
	case score <= 100:
		g = "a"
	}
	return g
}

func main()  {
	var n int
	fmt.Scanf("%d",&n)
	g:=grade(n)
	fmt.Println(g)
}
