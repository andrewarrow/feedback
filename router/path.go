package router

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/andrewarrow/feedback/files"
)

type Vars struct {
	Title  string
	Header string
	Footer string
}

func NewVars() Vars {
	v := Vars{}
	v.Title = "Feedback"
	return v
}

func NewVarsWithHeaderFooter() Vars {
	vars := NewVars()
	vars.Header = TemplateAsString("_header")
	vars.Footer = TemplateAsString("_footer")
	return vars
}

func TemplateAsString(name string) string {
	t, _ := template.ParseFiles(fmt.Sprintf("views/%s.html", name))
	vars := NewVars()
	var buffer bytes.Buffer
	t.Execute(&buffer, vars)
	fmt.Println(name, len(buffer.String()))
	return buffer.String()
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	fmt.Println(path)
	if path == "/" {
		t, _ := template.ParseFiles("views/welcome.html")
		vars := NewVarsWithHeaderFooter()
		t.Execute(writer, vars)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer)
	} else if path == "/feedback/add" {
		fmt.Fprintf(writer, "ok")
	} else {
		match := r.Paths[path]
		if match == "" {
			writer.WriteHeader(404)
			t, _ := template.ParseFiles("views/404.html")
			vars := NewVarsWithHeaderFooter()
			t.Execute(writer, vars)
		} else {
			t, _ := template.ParseFiles(fmt.Sprintf("views%s.html", path))
			vars := NewVarsWithHeaderFooter()
			t.Execute(writer, vars)
		}
	}
}

func (r *Router) HandleAsset(path string, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/css")
	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))
	fmt.Fprintf(writer, matchFile)
}
