package router

import (
	"os"
)

func (c *Context) LookupUser(guid string) map[string]any {
	return c.Router.LookupUser(guid)
}

func (c *Context) LookupUserByToken(token string) map[string]any {
	return c.Router.LookupUserByToken(token)
}

func (r *Router) LookupUserByToken(token string) map[string]any {
	var user map[string]any
	ct := r.SelectOne("cookie_token", "where guid=$1", []any{token})
	if len(ct) == 0 {
		return user
	}
	return r.SelectOne("user", "where id=$1", []any{ct["user_id"]})
}

func (r *Router) LookupUser(guid string) map[string]any {
	if guid == "" {
		return nil
	}
	params := []any{guid}
	m := r.SelectOne("user", "where guid=$1", params)
	if len(m) == 0 {
		return nil
	}
	return m
}

func (c *Context) LookupUsername(username string) map[string]any {
	return c.Router.LookupUsername(username)
}

func (r *Router) LookupUsername(username string) map[string]any {
	if username == "" {
		return map[string]any{}
	}
	params := []any{username}
	m := r.SelectOne("user", "where username=$1", params)
	if len(m) == 0 {
		return nil
	}
	return m
}

func IsAdmin(user map[string]any) bool {
	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "*" {
		return true
	}
	return user["guid"] == adminUser
}

func afterCreateUser(c *Context, guid string) {
}
