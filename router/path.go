package router

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type LayoutVars struct {
	Title   string
	Phone   string
	Content template.HTML
}

func (r *Router) PlaceContentInLayoutVars(filename string, vars any) *LayoutVars {
	content := new(bytes.Buffer)
	r.Template.ExecuteTemplate(content, filename, vars)

	lvars := LayoutVars{}
	lvars.Title = "test"
	lvars.Phone = "test"
	lvars.Content = template.HTML(content.String())
	return &lvars
}

func (r *Router) SendContentInLayout(writer http.ResponseWriter,
	filename string, contentVars any, status int) {
	vars := r.PlaceContentInLayoutVars(filename, contentVars)
	writer.WriteHeader(status)
	r.Template.ExecuteTemplate(writer, "application_layout.html", vars)
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	fmt.Println(path)
	if path == "/" {
		r.SendContentInLayout(writer, "welcome.html", nil, 200)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer)
	} else if !strings.HasSuffix(path, "/") {
		http.Redirect(writer, request, fmt.Sprintf("%s/", path), 301)
	} else {
		tokens := strings.Split(path, "/")
		first := tokens[1]
		match := r.Paths[first]
		if match == nil {
			r.SendContentInLayout(writer, "404.html", nil, 404)
		} else {
			c := Context{}
			c.writer = writer
			c.request = request
			c.router = r
			c.tokens = tokens[2:]
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
	if method == "GET" && len(tokens) == 1 {
		c.Index(context)
	} else if method == "GET" && len(tokens) > 1 {
		c.Show(context, tokens[1])
	} else if method == "POST" {
		//fmt.Printf("%+v\n", request.Header)
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(request.Body)
		fmt.Println("POST", buffer.String())
		contentType := request.Header["Content-Type"]
		if len(contentType) == 0 || contentType[0] == "application/x-www-form-urlencoded" {
			c.Create(context, buffer.String())
		} else {
			c.CreateWithJson(context, buffer.String())
		}
	}
}
