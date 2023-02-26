package router

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/controller"
	"github.com/andrewarrow/feedback/files"
)

func (r *Router) NewVarsWithHeaderFooter() controller.Vars {
	vars := controller.NewVars()
	vars.Header = TemplateAsString("_header", vars)
	vars.Phone = r.Site.Phone
	vars.Footer = TemplateAsString("_footer", vars)
	return vars
}

func TemplateAsString(name string, vars controller.Vars) template.HTML {
	t, _ := template.ParseFiles(fmt.Sprintf("views/%s.html", name))
	var buffer bytes.Buffer
	t.Execute(&buffer, vars)
	fmt.Println(name, len(buffer.String()))
	return template.HTML(buffer.String())
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	fmt.Println(path)
	if path == "/" {
		t, _ := template.ParseFiles("views/welcome.html")
		vars := r.NewVarsWithHeaderFooter()
		t.Execute(writer, vars)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer)
	} else if path == "/feedback/add" {
		fmt.Fprintf(writer, "ok")
	} else {
		match := r.Paths[path]
		if match == nil {
			writer.WriteHeader(404)
			t, _ := template.ParseFiles("views/404.html")
			vars := r.NewVarsWithHeaderFooter()
			t.Execute(writer, vars)
		} else {
			vars := r.NewVarsWithHeaderFooter()
			match.HandlePath(writer, path, vars)
			//t, _ := template.ParseFiles(fmt.Sprintf("views%s.html", match.ViewName()))
			//t.Execute(writer, vars)
		}
	}
}

func (r *Router) HandleAsset(path string, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/css")
	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))
	fmt.Fprintf(writer, matchFile)
}
