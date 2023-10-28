package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/util"
	"github.com/jmoiron/sqlx"
)

var BuildTag string

type Context struct {
	Writer       http.ResponseWriter
	Request      *http.Request
	tokens       []string
	Router       *Router
	User         map[string]any
	UserRequired bool
	path         string
	Db           *sqlx.DB
	NotFound     bool
	Method       string
	flash        string
	Layout       string
	Params       map[string]any
	Title        string
	LayoutMap    map[string]any
	ParamMutex   sync.Mutex
	Client       *http.Client
	Batch        bool
	BatchThing   any
}

func (c *Context) SendContentInLayout(filename string, vars any, status int) {
	if c.Title == "" {
		c.LayoutMap["title"] = c.Router.Site.Title
	} else {
		c.LayoutMap["title"] = models.RemoveMostNonAlphanumeric(c.Title)
	}
	c.LayoutMap["build"] = BuildTag
	ae := c.Request.Header.Get("Accept-Encoding")
	gzip := false
	if strings.Contains(ae, "gzip") {
		gzip = true
	}
	if c.Request.Header.Get("Feedback-Ajax") == "true" {
		c.SendContentForAjax(gzip, c.User, c.Writer, filename, vars, status)
		return
	}
	c.Router.SendContentInLayout(gzip, c.Layout, c.LayoutMap, c.flash, c.User, c.Writer, filename, vars, status)
}

func (c *Context) saveSchema() {
	asBytes, _ := json.Marshal(c.Router.Site)
	jqed := util.PipeToJq(string(asBytes))
	files.SaveFile("feedback.json", jqed)
}

func (c *Context) BodyAsString() string {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(c.Request.Body)
	c.Request.Body.Close()
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

	if c.Router.AllPosts && c.Method == "POST" && strings.HasPrefix(first, "Post") {
		funcToRun := c.Router.pathFuncToRun("posts")
		if len(tokens) == 3 {
			funcToRun(c, first, "")
		} else if len(tokens) == 4 {
			funcToRun(c, first, tokens[2])
		}
	}
	if c.Router.AllPosts && c.Method == "GET" && strings.HasPrefix(first, "Get") {
		funcToRun := c.Router.pathFuncToRun("gets")
		if len(tokens) == 3 {
			funcToRun(c, first, "")
		} else if len(tokens) == 4 {
			funcToRun(c, first, tokens[2])
		}
	}

	funcToRun := c.Router.pathFuncToRun(first)

	if funcToRun == nil {
		c.NotFound = true
		return
	}

	if len(tokens) == 3 { //          /foo/
		funcToRun(c, "", "")
	} else if len(tokens) == 4 { //   /foo/bar/
		funcToRun(c, tokens[2], "")
	} else if len(tokens) >= 5 { //   /foo/bar/more/
		funcToRun(c, tokens[2], tokens[3])
	} else {
		c.NotFound = true
	}
}
