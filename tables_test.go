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

func ExampleDropTable()  {
	table := NewTable("TestTable")
	status, err := DropTable("mysql",table)
	fmt.Println(status, err)
	// Output:
	// true <nil>
}