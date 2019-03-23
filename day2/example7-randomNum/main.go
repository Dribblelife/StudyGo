package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<6 ;i++  {
		fmt.Println("The random number is:",rand.Intn(10))
	}
}
