package meta_data

import (
	"context"

	"github.com/vishnu-chegondi/client"
)

// This returns all the tables in the db schema
func GetTables() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db_client := client.GetClient(ctx)

}
