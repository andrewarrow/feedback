package persist

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func SqliteConnection() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "sqlite.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
