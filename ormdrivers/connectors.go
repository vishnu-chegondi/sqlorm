package ormdrivers

import "database/sql/driver"

/*
GetConnector is a factory function returns a connector for the given driverName.
*/
func GetConnector(driverName string) (driver.Connector, error) {
	var connector driver.Connector
	var err error
	err = nil

	if driverName == "mysql" || driverName == "mariadb" {
		connector, err = mysqlConnector()
	}
	return connector, err
}
