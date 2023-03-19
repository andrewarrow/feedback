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
	model := r.Site.FindModel("user")
	params := []any{guid}
	return r.SelectOneFrom(model, "where guid=$1", params)
}

func (c *Context) LookupUsername(username string) map[string]any {
	return c.router.LookupUsername(username)
}

func (r *Router) LookupUsername(username string) map[string]any {
	if username == "" {
		return map[string]any{}
	}
	model := r.Site.FindModel("user")
	params := []any{username}
	return r.SelectOneFrom(model, "where username=$1", params)
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
