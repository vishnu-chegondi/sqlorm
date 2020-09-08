package migrations

import (
	"context"
	"fmt"

	"github.com/codeamenity/sqlorm/utils"

	"github.com/codeamenity/sqlorm/client"
	"github.com/codeamenity/sqlorm/models"
)

var OrmTables []models.TableInterface

// Get All Tables and create migration files
func MakeMigrations(client client.Client) {

	if isTable := CheckTableExists(client); isTable {
		// Check if columns change their values
		fmt.Print(isTable)
	} else {
		// Add CreateTable function to add table
		fmt.Print("Add Table")
	}

}

func CheckTableExists(client client.Client) bool {
	// Check if all table exists
	var data []map[string]interface{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, table := range OrmTables {
		tableName := table.GetTableName()
		rows, err := client.GetTable(ctx, tableName)
		if err == nil {
			data, err = utils.RowsToMap(rows)
			if len(data) == 0 {
				return false
			}

		}
	}
	return true
}

func GetTableDifferences(client client.Client) {
	// Check if all table exists

	// Check if columns are modified in existing tables

}
