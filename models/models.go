package models

type TableInterface interface {
	GetTableName() string
}

type BaseTable struct {
	tableName string
}

func (table *BaseTable) GetTableName() string {
	return table.tableName
}

