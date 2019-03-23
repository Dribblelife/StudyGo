package main

import (
	"fmt"
	"time"
)

func main() {
	i:= 0
	go func() {

		for  {
			i++
			fmt.Printf("goroutine i=%d",i)
			time.Sleep(time.Second)
		}
	}()

	for   {
		i++
		fmt.Printf("main i=%d",i)
		time.Sleep(time.Second)
		if i==3 {
			break
		}
	}
}
