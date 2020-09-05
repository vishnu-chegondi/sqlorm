package main

import (
	"github.com/codeamenity/sqlorm/client"
)

func main() {

	client := new(client.MYSQL)
	client.User = "mysqlorm"
	client.Passwd = "MYSQLpassword1"
	client.DBName = "mysqlorm"
	client.Net = "tcp"
	client.Addr = "localhost:3306"

	db, _ := client.GetDb()
	db.Ping()
}
