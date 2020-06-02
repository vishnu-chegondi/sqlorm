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
	table := NewTable("TestTable")
	table.Columns = []*Field{
		NewField("TestField", VarChar, "TestValue", true, 30),
	}
	status, err := CreateTable("mysql", table)
	fmt.Println(status, err)
	// Output:
	// true <nil>
}

func ExampleAddColumn()  {
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
	// false Cannot rename the column instead create new column and drop existing
}

func ExampleDropColumn()  {
	table := NewTable("TestTable")
	column := NewField("OtherColumn", Float, 10.23, false, 0)
	status, err := DropColumn("mysql", table, column)
	fmt.Println(status, err)
	// Output:
	// true <nil>
}

func ExampleDropTable()  {
	table := NewTable("TestTable")
	status, err := DropTable("mysql",table)
	fmt.Println(status, err)
	// Output:
	// true <nil>
}
