package client

import (
	"database/sql"
	"fmt"
	"os"

	mysql_connector "github.com/codeamenity/mysql_ormdriver/connector"
)

type MYSQL struct {
	User   string // Username
	Passwd string // Password (requires User)
	Net    string // Network type
	Addr   string // Network address (requires Net)
	DBName string // Database name
}

func (mysql *MYSQL) GetDb() (*sql.DB, error) {
	if mysql.User != "" {
		os.Setenv("User", mysql.User)
	}
	if mysql.Passwd != "" {
		os.Setenv("Passwd", mysql.Passwd)
	}
	if mysql.Net != "" {
		os.Setenv("Net", mysql.Net)
	}
	if mysql.Addr != "" {
		os.Setenv("Addr", mysql.Addr)
	}
	if mysql.DBName != "" {
		os.Setenv("DBName", mysql.DBName)
	}
	db, err := mysql_connector.GetDb()
	fmt.Printf("%v", db)
	return db, err
}
