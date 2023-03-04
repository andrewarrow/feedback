package router

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
)

type LayoutVars struct {
	Title   string
	User    *models.User
	Phone   string
	Content template.HTML
}

func (r *Router) PlaceContentInLayoutVars(user *models.User, filename string, vars any) *LayoutVars {
	content := new(bytes.Buffer)
	r.Template.ExecuteTemplate(content, filename, vars)

	lvars := LayoutVars{}
	lvars.Title = "Feedback"
	lvars.Phone = r.Site.Phone
	lvars.User = user
	lvars.Content = template.HTML(content.String())
	return &lvars
}

func (r *Router) SendContentInLayout(user *models.User, writer http.ResponseWriter,
	filename string, contentVars any, status int) {
	vars := r.PlaceContentInLayoutVars(user, filename, contentVars)
	writer.WriteHeader(status)
	r.Template.ExecuteTemplate(writer, "application_layout.html", vars)
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	cookie, err := request.Cookie("user")
	var user *models.User
	if err == nil && cookie.Value != "" {
		// use cookie.Value
		user = &models.User{}
		user.Username = "fred"
	}
	if path == "/" {
		r.SendContentInLayout(user, writer, "welcome.html", nil, 200)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer)
	} else if !strings.HasSuffix(path, "/") {
		http.Redirect(writer, request, fmt.Sprintf("%s/", path), 301)
	} else {
		c := Context{}
		c.writer = writer
		c.request = request
		c.router = r
		c.user = user
		c.path = path
		c.db = r.Db
		c.tokens = strings.Split(path, "/")
		c.userRequired = r.IsUserRequired(path, request.Method)
		handleContext(&c)
		if c.notFound {
			r.SendContentInLayout(user, writer, "404.html", nil, 404)
		}
	}
}

func handleModelsIndex(c *Context) {
	vars := ModelsVars{}
	vars.Models = c.router.Site.Models
	c.SendContentInLayout("models_index.html", vars, 200)
}

func (r *Router) HandleController(c Controller, context *Context) {
	//writer := c.context.writer
	request := context.request
	method := request.Method
	tokens := context.tokens
	//fmt.Println(method, context.path)
	if method == "GET" && len(tokens) == 1 {
		c.Index(context)
	} else if method == "GET" && len(tokens) > 1 {
		id := tokens[0]
		if id == "new" {
			c.New(context)
		} else {
			c.Show(context, id)
		}
	} else if method == "POST" {
		//fmt.Printf("%+v\n", request.Header)
		if context.userRequired && context.user == nil {
			http.Redirect(context.writer, context.request, "/sessions/new/", 302)
			return
		}
		contentType := request.Header["Content-Type"]
		if len(contentType) == 0 || contentType[0] == "application/x-www-form-urlencoded" {
			context.request.ParseForm()
			hiddenMethod := request.FormValue("_method")

			if hiddenMethod == "DELETE" {
				c.Destroy(context)
			} else {
				if len(tokens) == 1 {
					c.Create(context)
				} else {
					c.CreateWithId(context, tokens[0])
				}
			}
		} else {
			buffer := new(bytes.Buffer)
			buffer.ReadFrom(request.Body)
			fmt.Println("POST", buffer.String())
			c.CreateWithJson(context, buffer.String())
		}
	}
}
