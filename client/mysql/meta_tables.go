package mysql

import (
	"context"
	"database/sql"

	mysql_statements "github.com/codeamenity/mysql_ormdriver/statements"
)

type MetaTable struct {
}

func (db_driver *MYSQL) GetTable(ctx context.Context, tableName string) (*sql.Rows, error) {
	var rows *sql.Rows
	stmt, err := mysql_statements.GetTableStmnt(ctx, db_driver.Db)
	if err == nil {
		rows, err = stmt.QueryContext(ctx, tableName)
		defer stmt.Close()
	}
	return rows, err
}
