package router

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/models"
)

type Router struct {
	Paths             map[string]func() Controller
	UserRequiredPaths map[string]bool
	Template          *template.Template
	Site              *Site
}

type Context struct {
	writer       http.ResponseWriter
	request      *http.Request
	tokens       []string
	router       *Router
	user         *models.User
	userRequired bool
}

func (c *Context) SendContentInLayout(filename string, vars any, status int) {
	if c.userRequired && c.user == nil {
		http.Redirect(c.writer, c.request, "/sessions/new/", 301)
		return
	}
	c.router.SendContentInLayout(c.user, c.writer, filename, vars, status)
}

type Controller interface {
	Index(*Context)
	New(*Context)
	Create(*Context, string)
	CreateWithJson(*Context, string)
	Show(*Context, string)
	Destroy(*Context)
}

func NewRouter(path string) *Router {
	r := Router{}
	r.Paths = map[string]func() Controller{}
	r.UserRequiredPaths = map[string]bool{}

	var site Site
	jsonString := files.ReadFile(path)
	json.Unmarshal([]byte(jsonString), &site)
	r.Site = &site

	r.Template = LoadTemplates()
	r.Paths["models"] = NewModelsController
	r.Paths["sessions"] = NewSessionsController

	r.UserRequiredPaths["/sessions/new/"] = true
	r.UserRequiredPaths["/models/"] = true

	return &r
}
