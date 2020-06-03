package sqlorm

import (
	"database/sql"
	"github.com/codeamenity/sqlorm/ormdrivers"
)

func GetDb() *sql.DB {
	var db *sql.DB
	if ormdrivers.DriverName == "mysql" || ormdrivers.DriverName == "mariadb" {
		connector, err := ormdrivers.GetConnector(ormdrivers.DriverName)
		if err != nil {
			panic(err.Error())
		}
		db = sql.OpenDB(connector)
	}
	return db
}
