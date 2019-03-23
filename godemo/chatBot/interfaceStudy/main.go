package main

import "strings"

type Chatbot interface{
	Name() string
	Begin() (string,error)
	Talk
	ReportError(err error) string
	End() error
}
type Talk interface {
	Hello(userName string) string
	Talk(heard string)(saying string,end bool,err error)
}
type myTalk string

func (talk *myTalk) Hello(userName string) string  {

}

func (talk *myTalk) Talk(heard string) (saying string,end bool,err error)  {
	var talk Talk  = new(myTalk)
}

func (Chatbot *myTalk) Name() string  {

}

func (Chatbot *myTalk) Begin() (string,error)  {
	strings.TrimSpace()
}

