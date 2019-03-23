package main

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type Person struct {
	UserId int `db:"user_id"`
	UserName string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}

var DB *sqlx.DB

func init()  {
	database,err:=sqlx.Open("mysql","root:Mac_2019@tcp(localhost:3306)/test")
	if err!=nil {
		fmt.Println("open sql failed",err)
		return
	}
	DB=database
}

func main() {
	var person []Person
	_,err:=DB.Exec("update person set username=? where user_id=?","stu02",2)
	if err!=nil {
		fmt.Println("exec failed",err)
		return
	}
	fmt.Println("update succ",person)
}