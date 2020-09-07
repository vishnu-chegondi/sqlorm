package main

import (
	"github.com/codeamenity/sqlorm/client"
	"github.com/codeamenity/sqlorm/migrations"
	"github.com/codeamenity/sqlorm/models"
)

func main() {
	client := client.MYSQL{}
	client.User = "mysqlorm"
	client.Passwd = "MYSQLpassword1"
	client.DBName = "mysqlorm"
	client.Net = "tcp"
	client.Addr = "localhost:3306"
	client.Db = client.GetDb()

	TestTable := new(models.BaseTable)

	migrations.OrmTables = append(migrations.OrmTables, TestTable)
	// err := client.Ping()
	// if err != nil {
	// 	fmt.Print(err)
	// }

	migrations.MakeMigrations(client)
}
