package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	m:=make(map[string]interface{},4)
	m["company"]="Google"
	m["subjects"]=[]string{"go","java","c"}
	m["isOK"]=true
	m["price"]=66
	buf,err:=json.MarshalIndent(m,"","		")
	if err!=nil{
		fmt.Println(err)
		return
	}else {
		fmt.Println(string(buf))
	}
}
