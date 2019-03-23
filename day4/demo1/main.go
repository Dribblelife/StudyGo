package main

import "fmt"

func main() {
	var name  = "name1"
	fmt.Println(name)
	if true {
		name:="name2"
		varRename(name)
	}
}
func varRename(demoName string)  {
	fmt.Println(demoName)
}