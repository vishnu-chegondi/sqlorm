package client

import (
	"context"
	"database/sql"
	"os"

	mysql_connector "github.com/codeamenity/mysql_ormdriver/connector"
	mysql_statements "github.com/codeamenity/mysql_ormdriver/statements"
)

type Client interface {
	getDb() (*sql.DB, error)
	Ping() error
	GetAllTables() (*sql.Rows, error)
	GetTableColumns() (*sql.Rows, error)
}

type MYSQL struct {
	User   string // Username
	Passwd string // Password (requires User)
	Net    string // Network type
	Addr   string // Network address (requires Net)
	DBName string // Database name
	db     *sql.DB
}

func (db_driver *MYSQL) getDb() error {
	var err error
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
	if err = db_driver.db.Ping(); err != nil {
		db, err := mysql_connector.GetDb()
		if err == nil {
			db_driver.db = db
		}
	}
	return err
}

func (db_driver *MYSQL) Ping() error {
	err := db_driver.getDb()
	if err == nil {
		err = db_driver.db.Ping()
	}
	return err
}

func (db_driver *MYSQL) GetAllTables(ctx context.Context, DBName string) (*sql.Rows, error) {
	var rows *sql.Rows
	err := db_driver.getDb()
	stmt, err := mysql_statements.GetAllTablesStmnt(ctx, db_driver.db)
	if err == nil {
		rows, err = stmt.QueryContext(ctx, DBName)
		defer stmt.Close()
	}
	return rows, err
}

func (db_driver *MYSQL) GetTableColumns(ctx context.Context, DBName, tableName string) (*sql.Rows, error) {
	var rows *sql.Rows
	err := db_driver.getDb()
	stmt, err := mysql_statements.GetColumnsOfTable(ctx, db_driver.db)
	if err == nil {
		rows, err = stmt.QueryContext(ctx, DBName, tableName)
		defer stmt.Close()
	}
	return rows, err
}
