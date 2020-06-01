package statements

import (
	"database/sql"
	"fmt"
)

/*
CreateTableStmnt is used for generating a statement for creating a table of given database
*/
func CreateTableStmnt(db *sql.DB, database string) {
	fmt.Print("Test")
}
