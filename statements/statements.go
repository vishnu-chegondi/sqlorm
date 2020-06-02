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

/*
AddColumnStmnt returns the query string for adding the column
*/
func AddColumnStmnt(tableName string, driverName string) string {
	query := fmt.Sprintf("ALTER TABLE %s ADD", tableName)
	return query
}

/*
DropColumnStmnt returns the query string for droping the column
*/
func DropColumnStmnt(tableName string, driverName string) string {
	query := fmt.Sprintf("ALTER TABLE %s DROP COLUMN", tableName)
	return query
}

/*
RenameColumnStmnt returns the query string for renaming the column
*/
func RenameColumnStmnt(tableName string, driverName string) string {
	query := ""
	if driverName == "postgres"{
		query = fmt.Sprintf("ALTER TABLE %s RENAME COLUMN", tableName)
	}
	return query
}
