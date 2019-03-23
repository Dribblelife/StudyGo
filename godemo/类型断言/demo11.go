package main

import "fmt"

var container = []string{"zero","one","two"}

func main() {
	container := map[int]string{0:"zero",1:"one",2:"two"}
	fmt.Printf("The element is %q\n",container[1])
	//类型断言方式1
	_,ok1:=interface{}(container).([]string)
	_,ok2:=interface{}(container).(map[int]string)
	if !(ok1||ok2){
		fmt.Printf("error:unsupport container type:	%T\n",container)
		return
	}
	fmt.Printf("the elememt is %q.(container type: %T)\n",container[1],container)
	//方式2
	elem,err:=getElement(container)
	if err!=nil {
		fmt.Printf("Error is %s\n",err)
		return
	}
	fmt.Printf("the elememt is %q.(container type: %T)\n",elem,container)
}

func getElement(containerI interface{})(elem string,err error){
	switch t:=containerI.(type) {
	case []string:
		elem=t[1]
	case map[int]string:
		elem=t[1]
	default:
		err=fmt.Errorf("unsupport container type:	%T\n",containerI)
	}
	return
}
