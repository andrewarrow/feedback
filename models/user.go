package models

import "fmt"
import "github.com/jmoiron/sqlx"

type User struct {
	Id        int
	Email     string
	CreatedAt int64
}

func SelectUsers(db *sqlx.DB) ([]User, string) {
	users := []User{}
	sql := fmt.Sprintf("SELECT id, email, UNIX_TIMESTAMP(created_at) as createdat from users order by created_at desc")
	err := db.Select(&users, sql)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return users, s
}
