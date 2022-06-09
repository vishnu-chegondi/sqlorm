package mysql

import (
	"database/sql"

	mysql_connector "github.com/codeamenity/mysql_ormdriver/connector"
)

type Client interface {
	Ping() error
}

type MYSQL struct {
	Db *sql.DB
}

func NewMYSQL() *MYSQL {
	if db, err := mysql_connector.GetDb(); err == nil {
		mysql := &MYSQL{
			Db: db,
		}
		return mysql
	} else {
		panic(err.Error())
	}
}

func NewClient() Client {
	client := &MYSQL{}
	return client
}

func (db_driver *MYSQL) Ping() error {
	err := db_driver.Db.Ping()
	return err
}
