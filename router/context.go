package router

import (
	"bytes"
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
	method       string
}

func (c *Context) SendContentInLayout(filename string, vars any, status int) {
	c.router.SendContentInLayout(c.user, c.writer, filename, vars, status)
}

func (c *Context) BodyAsString() string {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(c.request.Body)
	return buffer.String()
}

func (c *Context) ReadFormPost() {
	c.request.ParseForm()
	c.method = c.request.Method
	hiddenMethod := c.request.FormValue("_method")
	if hiddenMethod != "" {
		c.method = hiddenMethod
	}
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