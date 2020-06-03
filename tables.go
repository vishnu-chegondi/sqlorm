package sqlorm

import (
	"errors"
	"fmt"
	"strings"

	"github.com/codeamenity/sqlorm/ormdrivers"
	"github.com/codeamenity/sqlorm/statements"
)

type Table struct {
	Name       string
	Columns    []*Field
	PrimaryKey string
}

/*
NewTable creates a new table from Table Struct with given name as table name
*/
func NewTable(Name string) *Table {
	table := new(Table)
	table.Name = Name
	return table
}

func (tb *Table) getPrimaryStmnt() string {
	var stmnt string
	if tb.PrimaryKey != "" {
		stmnt = fmt.Sprintf("PRIMARY KEY (%s)", tb.PrimaryKey)
	} else {
		stmnt = fmt.Sprint("Id int NOT NULL AUTO_INCREMENT PRIMARY KEY")
	}
	return stmnt
}

/*
CreateTable is used for creating a table in the database.
If no columns are provided then default Id(Column) as primary key is created
*/
func CreateTable(tb *Table) (bool, error) {
	var colStmnt []string
	db := GetDb()
	defer db.Close()
	create_query := statements.CreateTableStmnt(tb.Name, ormdrivers.DriverName)
	for _, column := range tb.Columns {
		if column.ValidType() && column.ValidDefaultValue() {
			colStmnt = append(colStmnt, column.GetColumnStmnt())
		}
	}
	colStmnt = append(colStmnt, tb.getPrimaryStmnt())
	col_string := strings.Join(colStmnt, ",")
	create_query += col_string + ")"
	_, err := db.Query(create_query)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
DropTable is used for dropping a table in the database.
*/
func DropTable(tb *Table) (bool, error) {
	db := GetDb()
	defer db.Close()
	drop_query := statements.DropTableStmnt(tb.Name, ormdrivers.DriverName)
	_, err := db.Query(drop_query)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
AddColumns is used for adding columns to existing tables with default values
*/
func AddColumn(tb *Table, column *Field) (bool, error) {
	db := GetDb()
	defer db.Close()
	add_column := statements.AddColumnStmnt(tb.Name, ormdrivers.DriverName)
	add_column += " " + column.GetColumnStmnt()
	_, err := db.Query(add_column)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
DropColumn is used for removing columns in the table.
*/
func DropColumn(tb *Table, column *Field) (bool, error) {
	db := GetDb()
	defer db.Close()
	drop_column := statements.DropColumnStmnt(tb.Name, ormdrivers.DriverName)
	drop_column += " " + column.Name
	_, err := db.Query(drop_column)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
RenameColumn is used for renaming column with out any data loss. This is not supported on every database on every versions.

Example: This only works on mysql 8.0
*/
func RenameColumn(tb *Table, column *Field, newName string) (bool, error) {
	db := GetDb()
	defer db.Close()
	rename_column := statements.RenameColumnStmnt(tb.Name, ormdrivers.DriverName)
	if rename_column == "" {
		return false, errors.New("Cannot rename the column instead use ChangeColumn")
	}
	rename_column = fmt.Sprintf("%s %s TO %s", rename_column, column.Name, newName)
	_, err := db.Query(rename_column)
	if err != nil {
		return false, err
	}
	return true, nil
}

/*
ChangeColumn is used for changing the column type definition and even name
*/
func ChangeColumn(tb *Table, column *Field, newcolumn *Field) (bool, error) {
	db := GetDb()
	defer db.Close()
	change_column := statements.ChangeColumnStmnt(tb.Name, ormdrivers.DriverName, column.Name, newcolumn.GetColumnStmnt())
	_, err := db.Query(change_column)
	if err != nil {
		return false, err
	}
	return true, nil
}
