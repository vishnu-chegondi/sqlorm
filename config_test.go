package sqlorm

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	msg := make(chan string)
	db := GetDb("mysql")
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	go DBPing(db, msg)
	select {
	case text := <-msg:
		fmt.Print(text)
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
