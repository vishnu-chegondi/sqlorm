package utils

import (
	"database/sql"
	"reflect"
)

func RowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	var err error

	cols, err := rows.Columns()
	if err != nil {
		return data, err
	}

	for rows.Next() {
		// array of different types and pointer to indicate them
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))

		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return data, err
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			if reflect.TypeOf(columns[i]) == reflect.TypeOf([]byte{}) {
				val := columns[i].([]byte)
				m[colName] = string(val)
			} else if columns[i] == nil {
				m[colName] = nil
			} else {
				val := columns[i].(interface{})
				m[colName] = val
			}
		}

		data = append(data, m)
	}
	return data, err
}
