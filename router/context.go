package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Context struct {
	Writer       http.ResponseWriter
	Request      *http.Request
	tokens       []string
	router       *Router
	User         *User
	UserRequired bool
	path         string
	Db           *sqlx.DB
	NotFound     bool
	Method       string
	flash        string
	Title        string
	Layout       string
}

func (c *Context) SendContentInLayout(filename string, vars any, status int) {
	if c.Title == "" {
		c.Title = c.router.Site.Title
	}
	c.router.SendContentInLayout(c.Layout, c.Title, c.flash, c.User, c.Writer, filename, vars, status)
}

func (c *Context) saveSchema() {
	asBytes, _ := json.Marshal(c.router.Site)
	c.Db.Exec(fmt.Sprintf("update feedback_schema set json_string = '%s'", string(asBytes)))
}

func (c *Context) BodyAsString() string {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(c.Request.Body)
	return buffer.String()
}

func (c *Context) ReadFormPost() {
	hiddenMethod := c.Request.FormValue("_method")
	if hiddenMethod != "" {
		c.Method = hiddenMethod
	}
}

func handleContext(c *Context) {
	tokens := c.tokens
	first := tokens[1]
	funcToRun := c.router.Paths[first]

	if funcToRun == nil {
		c.NotFound = true
		return
	}

	if len(tokens) == 3 { //          /foo/
		funcToRun(c, "", "")
	} else if len(tokens) == 4 { //   /foo/bar/
		funcToRun(c, tokens[2], "")
	} else if len(tokens) == 5 { //   /foo/bar/more/
		funcToRun(c, tokens[2], tokens[3])
	} else {
		c.NotFound = true
	}
}
