package main

import (
	"fmt"

	"github.com/codeamenity/sqlorm/client"
	"github.com/codeamenity/sqlorm/migrations"
	"github.com/codeamenity/sqlorm/models"
)

func main() {
	client := &client.MYSQL{
		User:   "mysqlorm",
		Passwd: "MYSQLpassword1",
		DBName: "mysqlorm",
		Net:    "tcp",
		Addr:   "localhost:3306",
	}

	TestTable := new(models.BaseTable)

	migrations.OrmTables = append(migrations.OrmTables, TestTable)
	err := client.Ping()
	if err != nil {
		fmt.Print(err)
	}

	migrations.MakeMigrations(client)
}
