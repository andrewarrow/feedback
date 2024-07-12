package persist

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func SqliteConnection(name string) *sqlx.DB {
	currentUser, _ := user.Current()
	prefix := currentUser.HomeDir + "/" + name
	if strings.HasPrefix(name, "/") {
		prefix = name
	}
	db, err := sqlx.Connect("sqlite3", prefix+"_sqlite_560dc8c4-b18a-4517-a90c-b0f92d2ba5a5.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
