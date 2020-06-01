package sqlorm

import "fmt"

func ExampleField()  {
	field := NewField("test", VarChar, "TestValue")
	if field.validType() &&	field.ValidDefaultValue() {
		fmt.Println("Successfully created field")
	}
}