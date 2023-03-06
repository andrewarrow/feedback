package router

import (
	"fmt"
	"os"
)

type User struct {
	Username  string
	Timestamp string
	Ago       string
	Guid      string
}

func (r *Router) LookupUser(guid string) *User {
	if guid == "" {
		return nil
	}
	rows, err := r.Db.Queryx("SELECT * FROM users where guid=$1", guid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	m := make(map[string]any)
	rows.Next()
	rows.MapScan(m)
	if len(m) == 0 {
		return nil
	}
	user := User{}
	user.Username = fmt.Sprintf("%s", m["username"])
	user.Guid = guid
	return &user
}

func (r *Router) LookupUsername(username string) *User {
	if username == "" {
		return nil
	}
	rows, err := r.Db.Queryx("SELECT * FROM users where username=$1", username)
	if err != nil {
		return nil
	}
	defer rows.Close()
	m := make(map[string]any)
	rows.Next()
	rows.MapScan(m)
	if len(m) == 0 {
		return nil
	}
	user := User{}
	user.Username = fmt.Sprintf("%s", m["username"])
	user.Timestamp, user.Ago = FixTime(m)
	return &user
}

func (u *User) IsAdmin() bool {
	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "*" {
		return true
	}
	return u.Guid == adminUser
}
