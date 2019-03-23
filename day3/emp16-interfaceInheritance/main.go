package main

import "fmt"

type humaner interface {
	sayHi()
}

type personer interface {
	humaner
	sing(irc string)
}

type student struct {
	name string
	id int
}

func (tmp *student) sayHi(){
	fmt.Printf("Student[%s %d] sayhi\n",tmp.name,tmp.id)

}

func (tmp *student) sing(irc string) {
	fmt.Printf("Student[%s %d]唱:%s",tmp.name,tmp.id,irc)
}

func main() {
	var i personer
	s:=&student{"mac",1}
	i=s
	i.sayHi()
	i.sing("起来！")

}
