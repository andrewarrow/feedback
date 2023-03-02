package router

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/andrewarrow/feedback/controller"
	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/models"
)

type Router struct {
	Paths    map[string]func() Controller
	Template *template.Template
	Vars     *controller.Vars
	Site     *controller.Site
}

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
	tokens  []string
	router  *Router
	user    *models.User
}

func (c *Context) SendContentInLayout(filename string, vars any, status int) {
	c.router.SendContentInLayout(c.user, c.writer, filename, vars, status)
}

type Controller interface {
	Index(*Context)
	New(*Context)
	Create(*Context, string)
	CreateWithJson(*Context, string)
	Show(*Context, string)
}

func NewRouter(path string) *Router {
	r := Router{}
	r.Paths = map[string]func() Controller{}

	var site controller.Site
	jsonString := files.ReadFile(path)
	json.Unmarshal([]byte(jsonString), &site)
	r.Site = &site

	r.Vars = controller.NewVars(&site)
	r.Template = LoadTemplates()
	r.Paths["models"] = NewModelsController
	r.Paths["sessions"] = NewSessionsController

	return &r
}
