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
		tokens := strings.Split(path, "/")
		first := tokens[1]
		match := r.Paths[first]
		if match == nil {
			r.SendContentInLayout(user, writer, "404.html", nil, 404)
		} else {
			c := Context{}
			c.writer = writer
			c.request = request
			c.router = r
			c.user = user
			c.path = path
			c.tokens = tokens[2:]
			c.userRequired = r.IsUserRequired(path, request.Method)
			controller := match()
			r.HandleController(controller, &c)
		}
	}
}

func (r *Router) HandleController(c Controller, context *Context) {
	//writer := c.context.writer
	request := context.request
	method := request.Method
	tokens := context.tokens
	fmt.Println(method, context.path)
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
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(request.Body)
		fmt.Println("POST", buffer.String())
		contentType := request.Header["Content-Type"]
		if len(contentType) == 0 || contentType[0] == "application/x-www-form-urlencoded" {
			payload := buffer.String()
			if payload == "_method=DELETE" {
				c.Destroy(context)
			} else {
				c.Create(context, payload)
			}
		} else {
			c.CreateWithJson(context, buffer.String())
		}
	}
}
