package sqlorm

import (
	"database/sql"
	"github.com/vishnu-chegondi/sqlorm/ormdrivers"
)

func GetDb(driverName string) *sql.DB {
	var db *sql.DB
	if driverName == "mysql" || driverName == "mariadb" {
		connector, err := ormdrivers.GetConnector(driverName)
		if err != nil {
			panic(err.Error())
		}
		db = sql.OpenDB(connector)
	}
	return db
}
