package sqlorm

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/codeamenity/sqlorm/ormdrivers"
)

func TestMain(m *testing.M) {
	msg := make(chan string)
	ormdrivers.DBName = "mysqlorm"
	db := GetDb()
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	go DBPing(db, msg)
	select {
	case text := <-msg:
		fmt.Println(text)
		os.Exit(m.Run())
	case <-ctx.Done():
		ctx.Err()
		fmt.Print("Unable to ping db")
		os.Exit(2)
	}
}

func DBPing(db *sql.DB, msg chan string) {
	error := db.Ping()
	if error != nil {
		time.Sleep(10 * time.Second)
		DBPing(db, msg)
	} else {
		msg <- "success"
	}
}

func ExampleGetDb() {
	ormdrivers.Addr = "localhost:3306"
	db := GetDb()
	defer db.Close()
	err := db.Ping()
	fmt.Println(err)
	// Output:
	// <nil>
}
