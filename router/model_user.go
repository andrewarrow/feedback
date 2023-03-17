package router

import (
	"os"
)

type User struct {
	Username  string
	Timestamp string
	Ago       string
	Guid      string
	Id        int64
}

func (r *Router) LookupUser(guid string) *User {
	if guid == "" {
		return nil
	}
	model := r.Site.FindModel("user")
	params := []any{guid}
	m := r.SelectOneFrom(model, "where guid=$1", params)

	if len(m) == 0 {
		return nil
	}
	user := User{}
	user.Username = m["username"].(string)
	user.Guid = guid
	user.Id = m["id"].(int64)
	return &user
}

func (r *Router) LookupUsername(username string) *User {
	if username == "" {
		return nil
	}
	model := r.Site.FindModel("user")
	params := []any{username}
	m := r.SelectOneFrom(model, "where username=$1", params)

	if len(m) == 0 {
		return nil
	}
	user := User{}
	user.Username = m["username"].(string)
	user.Timestamp = m["created_at"].(string)
	user.Ago = m["created_at_ago"].(string)
	return &user
}

func (u *User) IsAdmin() bool {
	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "*" {
		return true
	}
	return u.Guid == adminUser
}
