package router

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/files"
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

func (r *Router) SendContentInLayout(filename string, contentVars any) {
	vars := r.PlaceContentInLayoutVars(filename, contentVars)
	r.Template.ExecuteTemplate(r.Writer, "application_layout.html", vars)
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	fmt.Println(path)
	r.Writer = writer
	if path == "/" {
		r.SendContentInLayout("welcome.html", nil)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer)
	} else if !strings.HasSuffix(path, "/") {
		http.Redirect(writer, request, fmt.Sprintf("%s/", path), 301)
	} else {
		tokens := strings.Split(path, "/")
		first := tokens[1]
		match := r.Paths[first]
		if match == nil {
			writer.WriteHeader(404)
			r.Template.ExecuteTemplate(writer, "404.html", r.Vars)
		} else {
			match.HandlePath(writer, request, tokens[2:])
		}
	}
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
