package statements

import (
	"fmt"
)

/*
CreateTableStmnt is used for generating a statment for creating a table of given database
*/
func CreateTableStmnt(tableName string, driverName string) string {
	query := fmt.Sprintf("CREATE TABLE %s (", tableName)
	return query
}

/*
DropTableStmnt is used for generating a statment for droping a table in given database
*/
func DropTableStmnt(tableName string, driverName string) string {
	query:= fmt.Sprintf("DROP TABLE %s", tableName)
	return query
}