package main

import (
		"fmt"
	"time"
	)

func main() {
	timer:=time.NewTicker(1*time.Second)
	i:=0
	for   {
		<-timer.C
		i++
		fmt.Println("i=",i)
		if i==5 {
			timer.Stop()
			break
		}
	}
}
