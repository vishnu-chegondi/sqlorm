package sqlorm

import (
	"fmt"
	"reflect"

	"github.com/codeamenity/sqlorm/ormdrivers"
)

var fieldTypes = []string{"VARCHAR", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT", "TINYINT", "MEDIUMINT", "INT", "BIGINT", "FLOAT", "DOUBLE", "DECIMAL", "DATE", "DATETIME", "TIMESTAMP", "TIME", "ENUM", "BOOLEAN"}

const (
	VarChar    = "VARCHAR"
	TinyText   = "TINYTEXT"
	Text       = "TEXT"
	MediumText = "MEDIUMTEXT"
	LongText   = "LONGTEXT"
	TinyInt    = "TINYINT"
	MediumInt  = "MEDIUMINT"
	Int        = "INT"
	BigInt     = "BIGINT"
	Float      = "FLOAT"
	Double     = "DOUBLE"
	Decimal    = "DECIMAL"
	Date       = "DATE"
	DateTime   = "DATETIME"
	TimeStamp  = "TIMESTAMP"
	Time       = "TIME"
	Enum       = "ENUM"
	Boolean    = "BOOLEAN"
)

type Field struct {
	Name           string
	Type           string
	DefaultValue   interface{}
	Required       bool
	MaxValue       int
	ForeignKey     bool
	ReferenceTable *Table
}

/*
NewField Returns a field with
*/
func NewField(Name string, FieldType string, DefaultValue interface{}, Required bool, MaxValue int) *Field {
	field := new(Field)
	field.Name = Name
	field.Type = FieldType
	field.Required = Required
	field.MaxValue = MaxValue
	field.DefaultValue = DefaultValue
	return field
}

/*
ValidType checks if the given type for a field is valid or not.
*/
func (field *Field) ValidType() bool {
	_, status := Find(field.Type, fieldTypes)
	if status != true {
		error_msg := fmt.Sprintf("Field %s is not given a valid type. Assigned Type is %s", field.Name, field.Type)
		panic(error_msg)
	}
	return true
}

/*
ValidDefaultValue checks if the given default value if of field type if not panics
*/
func (field *Field) ValidDefaultValue() bool {
	v := reflect.ValueOf(field.DefaultValue)
	error_msg := fmt.Sprintf("Field %s is not given a valid DefaultValue. Assigned DefaultValue is %s", field.Name, field.Type)
	if _, status := Find(field.Type, fieldTypes[:5]); status || field.Type == "DECIMAL" || field.Type == "ENUM" {
		if v.Kind() != reflect.String {
			panic(error_msg)
		}
	} else if field.Type == "TINYINT" {
		if v.Kind() != reflect.Int16 {
			panic(error_msg)
		}
	} else if field.Type == "MEDIUMINT" {
		if v.Kind() != reflect.Int32 {
			panic(error_msg)
		}
	} else if field.Type == "INT" {
		if v.Kind() != reflect.Int {
			panic(error_msg)
		}
	} else if field.Type == "BIGINT" {
		if v.Kind() != reflect.Int64 {
			panic(error_msg)
		}
	} else if field.Type == "FLOAT" {
		if v.Kind() != reflect.Float32 {
			panic(error_msg)
		}
	} else if field.Type == "DOUBLE" {
		if v.Kind() != reflect.Float64 {
			panic(error_msg)
		}
	} else if field.Type == "BOOLEAN" {
		if v.Kind() != reflect.Bool {
			panic(error_msg)
		}
	} else if _, status := Find(field.Type, fieldTypes[12:16]); status {
		if v.Kind() != reflect.Struct {
			panic(error_msg)
		}
	}
	return true
}

/*
GetColumnStmnt returns the statement required indicated by Field Passed depending on the driverName
*/
func (field *Field) GetColumnStmnt() string {
	var colStmnt string
	if ormdrivers.DriverName == "mysql" {
		colStmnt = fmt.Sprintf("%s %s", field.Name, field.Type)
		if field.MaxValue != 0 {
			colStmnt = fmt.Sprintf("%s %s(%d)", field.Name, field.Type, field.MaxValue)
		}
		if field.Required {
			colStmnt += " NOT NULL"
		}
		if field.ForeignKey {
			colStmnt += " " + field.GetForeingKeyStmnt()
		}
	}
	return colStmnt
}

/*
GetForeignKey returns the statement required for attaching foreignKey to the table
*/
func (field *Field) GetForeingKeyStmnt() string {
	var foreignStmnt string
	if ormdrivers.DriverName == "mysql" {
		foreignStmnt = fmt.Sprintf("FOREIGN KEY REFERENCES %s(%s)", field.ReferenceTable.Name, field.ReferenceTable.PrimaryKey)
	}
	return foreignStmnt
}
