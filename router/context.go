package router

import (
	"net/http"

	"github.com/andrewarrow/feedback/models"
	"github.com/jmoiron/sqlx"
)

type Context struct {
	writer       http.ResponseWriter
	request      *http.Request
	tokens       []string
	router       *Router
	user         *models.User
	userRequired bool
	path         string
	db           *sqlx.DB
	notFound     bool
}

func (c *Context) SendContentInLayout(filename string, vars any, status int) {
	if c.userRequired && c.user == nil {
		http.Redirect(c.writer, c.request, "/sessions/new/", 302)
		return
	}
	c.router.SendContentInLayout(c.user, c.writer, filename, vars, status)
}

func handleContext(c *Context) {
	tokens := c.tokens

	if len(tokens) == 3 { //          /foo/
		handlePathContext(c, tokens[1], "", "")
	} else if len(tokens) == 4 { //   /foo/bar/
		handlePathContext(c, tokens[1], tokens[2], "")
	} else if len(tokens) == 5 { //   /foo/bar/more/
		handlePathContext(c, tokens[1], tokens[2], tokens[3])
	} else {
		c.notFound = true
	}
}

func handlePathContext(c *Context, first, second, third string) {
	if first == "models" {
		handleModels(c, second, third)
	} else if first == "sessions" {
		handleSessions(c, second, third)
	} else {
		c.notFound = true
	}
}
