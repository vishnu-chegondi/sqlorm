package client

import (
	"context"
	"database/sql"
	"os"
	"sync"

	mysql_connector "github.com/codeamenity/mysql_ormdriver/connector"
	mysql_statements "github.com/codeamenity/mysql_ormdriver/statements"
)

type Client interface {
	setDb()
	Ping() error
	GetTable(ctx context.Context, tableName string) (*sql.Rows, error)
	GetTableColumns(ctx context.Context, tableName string) (*sql.Rows, error)
}

type MYSQL struct {
	User   string // Username
	Passwd string // Password (requires User)
	Net    string `default:"tcp"` // Network type
	Addr   string // Network address (requires Net)
	DBName string // Database name
	Db     *sql.DB
}

func (db_driver *MYSQL) setDb() {
	var once sync.Once
	once.Do(func() {
		if db_driver.User != "" {
			os.Setenv("User", db_driver.User)
		}
		if db_driver.Passwd != "" {
			os.Setenv("Passwd", db_driver.Passwd)
		}
		if db_driver.Net != "" {
			os.Setenv("Net", db_driver.Net)
		}
		if db_driver.Addr != "" {
			os.Setenv("Addr", db_driver.Addr)
		}
		if db_driver.DBName != "" {
			os.Setenv("DBName", db_driver.DBName)
		}
		db, err := mysql_connector.GetDb()
		if err != nil {
			panic(err.Error())
		}
		db_driver.Db = db
	})
}

func (db_driver *MYSQL) Ping() error {
	db_driver.setDb()
	err := db_driver.Db.Ping()
	return err
}

func (db_driver *MYSQL) GetTable(ctx context.Context, tableName string) (*sql.Rows, error) {
	db_driver.setDb()
	var rows *sql.Rows
	stmt, err := mysql_statements.GetTableStmnt(ctx, db_driver.Db)
	if err == nil {
		rows, err = stmt.QueryContext(ctx, db_driver.DBName, tableName)
		defer stmt.Close()
	}
	return rows, err
}

func (db_driver *MYSQL) GetTableColumns(ctx context.Context, tableName string) (*sql.Rows, error) {
	db_driver.setDb()
	var rows *sql.Rows
	stmt, err := mysql_statements.GetColumnsOfTable(ctx, db_driver.Db)
	if err == nil {
		rows, err = stmt.QueryContext(ctx, db_driver.DBName, tableName)
		defer stmt.Close()
	}
	return rows, err
}
