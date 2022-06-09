package models

import "reflect"

func NewTableField(Name string, ColumnType reflect.Type, DefaultValue interface{}) *TableField {
	return &TableField{
		Name:         Name,
		ColumnType:   ColumnType,
		DefaultValue: DefaultValue,
	}
}

type TableField struct {
	Name         string
	ColumnType   reflect.Type
	DefaultValue interface{}
}

// Validates if given default value is of type columnType
func (tablefield *TableField) ValidateDefaultValue() {
}
