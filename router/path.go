package router

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/models"
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
			match(writer)
		}
	}
}

type ModelsVars struct {
	Models []models.Model
}

func (r *Router) ModelsIndex(writer http.ResponseWriter) {
	vars := ModelsVars{}
	vars.Models = r.Site.Models
	r.SendContentInLayout(writer, "models_index.html", vars, 200)
}

func (r *Router) HandleAsset(path string, writer http.ResponseWriter) {
	contentType := "text/css"
	if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	}
	writer.Header().Set("Content-Type", contentType)
	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))
	fmt.Fprintf(writer, matchFile)
}
