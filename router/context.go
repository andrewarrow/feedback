package router

import (
	"net/http"

	"github.com/andrewarrow/feedback/models"
)

type Context struct {
	writer       http.ResponseWriter
	request      *http.Request
	tokens       []string
	router       *Router
	user         *models.User
	userRequired bool
	path         string
}

func (c *Context) SendContentInLayout(filename string, vars any, status int) {
	if c.userRequired && c.user == nil {
		http.Redirect(c.writer, c.request, "/sessions/new/", 302)
		return
	}
	c.router.SendContentInLayout(c.user, c.writer, filename, vars, status)
}
