package sqlorm

import (
	"fmt"
)

func ExampleNewTable() {
	table := NewTable("TestTable")
	fmt.Println(table.Name)
	// Output:
	// TestTable
}

func ExampleCreateTable() {
	foreignTable := NewTable("foreignTable")
	status, err := CreateTable("mysql", foreignTable)
	fmt.Println(status, err)
	foreignField := NewField("FoeignField", Int, 0, true, 0)
	foreignField.ForeignKey = true
	foreignField.ReferenceTable = foreignTable
	table := NewTable("TestTable")
	table.Columns = []*Field{
		NewField("TestField", VarChar, "TestValue", true, 30),
	}
	status, err = CreateTable("mysql", table)
	fmt.Println(status, err)
	// Output:
	// true <nil>
	// true <nil>
}

func ExampleAddColumn() {
	table := NewTable("TestTable")
	column := NewField("OtherColumn", Float, 10.23, false, 0)
	status, err := AddColumn("mysql", table, column)
	fmt.Println(status, err)
	// Output:
	// true <nil>
}

func ExampleRenameColumn() {
	table := NewTable("TestTable")
	column := NewField("OtherColumn", Float, 10.23, false, 0)
	status, err := RenameColumn("mysql", table, column, "NewName")
	fmt.Println(status, err)
	// Output:
	// false Cannot rename the column instead use ChangeColumn
}

func ExampleChangeColumn() {
	table := NewTable("TestTable")
	column := NewField("OtherColumn", Float, 10.23, false, 0)
	newcolumn := NewField("NewColumn", Double, 10.234, true, 0)
	status, err := ChangeColumn("mysql", table, column, newcolumn)
	fmt.Println(status, err)
	// Output:
	// true <nil>
}

func ExampleDropColumn() {
	table := NewTable("TestTable")
	column := NewField("NewColumn", Float, 10.23, false, 0)
	status, err := DropColumn("mysql", table, column)
	fmt.Println(status, err)
	// Output:
	// true <nil>
}

func ExampleDropTable() {
	foreignTable := NewTable("foreignTable")
	status, err := DropTable("mysql", foreignTable)
	fmt.Println(status, err)
	table := NewTable("TestTable")
	status, err = DropTable("mysql", table)
	fmt.Println(status, err)
	// Output:
	// true <nil>
	// true <nil>
}
