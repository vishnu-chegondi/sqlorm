package sqlorm

import "fmt"

func ExampleField() {
	field := NewField("test", VarChar, "TestValue", false, 0)
	if field.ValidType() && field.ValidDefaultValue() {
		fmt.Println("Successfully created field")
	}
	// Output:
	// Successfully created field
}
