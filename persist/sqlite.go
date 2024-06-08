package persist

import (
	"fmt"

	"github.com/andrewarrow/feedback/files"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func SqliteConnection(name string) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", files.UserHomeDir()+"/"+name+"_sqlite_560dc8c4-b18a-4517-a90c-b0f92d2ba5a5.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
