package sqlorm

import (
	"fmt"
	"strings"
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
func CreateTable(driverName string, tb *Table) (bool, error) {
	var colStmnt []string
	db := GetDb(driverName)
	create_query := statements.CreateTableStmnt(tb.Name, driverName)
	for _, column := range tb.Columns {
		if column.ValidType() && column.ValidDefaultValue() {
			colStmnt = append(colStmnt, column.GetColumnStmnt(driverName))
		}
	}
	colStmnt = append(colStmnt, tb.getPrimaryStmnt())
	col_string := strings.Join(colStmnt, ",")
	create_query += col_string+")"
	_, err := db.Query(create_query)
	if err!=nil{
		return false, err
	}
	return true, nil
}

/*
DropTable is used for dropping a table in the database.
*/
func DropTable(driverName string, tb *Table) (bool, error) {
	db := GetDb(driverName)
	drop_query := statements.DropTableStmnt(tb.Name, driverName)
	_, err := db.Query(drop_query)
	if err!=nil{
		return false, err
	}
	return true, nil
}