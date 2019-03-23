package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	n:=rand.Intn(100)
	for  {
		var input int
		fmt.Scanf("%d",&input)
		flag:=false
		switch  {
		case input ==n:
			fmt.Println("you are right")
			flag=true
		case input < n:
			fmt.Println("small")
		case input > n:
			fmt.Println("big")

		}
		if flag{
			break
		}
		
	}
}
