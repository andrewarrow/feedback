package router

import (
	"os"
)

func (c *Context) LookupUser(guid string) map[string]any {
	return c.router.LookupUser(guid)
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
	return c.router.LookupUsername(username)
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
