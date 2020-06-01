package sqlorm

import "fmt"

func ExampleField() {
	field := NewField("test", VarChar, "TestValue")
	if field.ValidType() && field.ValidDefaultValue() {
		fmt.Println("Successfully created field")
	}
	// Output:
	// Successfully created field
}
