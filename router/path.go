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
	Title    string
	SiteName string
	User     *User
	Footer   string
	Content  template.HTML
	Flash    string
}

func (r *Router) PlaceContentInLayoutVars(title, flash string, user *User, filename string, vars any) *LayoutVars {
	content := new(bytes.Buffer)
	r.Template.ExecuteTemplate(content, filename, vars)

	lvars := LayoutVars{}
	lvars.Title = models.RemoveMostNonAlphanumeric(title)
	lvars.Footer = r.Site.Footer
	lvars.SiteName = r.Site.Title
	lvars.User = user
	lvars.Flash = flash
	lvars.Content = template.HTML(content.String())
	return &lvars
}

func (r *Router) SendContentInLayout(layout, title, flash string, user *User, writer http.ResponseWriter,
	filename string, contentVars any, status int) {
	vars := r.PlaceContentInLayoutVars(title, flash, user, filename, contentVars)
	writer.WriteHeader(status)
	r.Template.ExecuteTemplate(writer, layout, vars)
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	cookie, err := request.Cookie("user")
	var user *User
	if err == nil && cookie.Value != "" {
		user = r.LookupUser(cookie.Value)
	}
	cookie, err = request.Cookie("flash")
	flash := ""
	if err == nil && cookie.Value != "" {
		flash = cookie.Value
		removeFlash(writer)
	}

	if path == "/" {
		funcToCall := r.Paths["/"]
		c := PrepareContext(r, user, "/", flash, writer, request)
		if funcToCall == nil {
			r.SendContentInLayout(c.Layout, r.Site.Title, flash, user, writer, "welcome.html",
				nil, 200)
		} else {
			funcToCall(c, "", "")
		}
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer)
	} else if !strings.HasSuffix(path, "/") {
		http.Redirect(writer, request, fmt.Sprintf("%s/", path), 301)
	} else {
		c := PrepareContext(r, user, path, flash, writer, request)
		c.tokens = strings.Split(path, "/")
		if c.Method == "POST" {
			c.ReadFormPost()
		}
		handleContext(c)
		if c.UserRequired && c.User == nil {
			http.Redirect(c.Writer, c.Request, "/sessions/new/", 302)
			return
		}
		if c.NotFound {
			r.SendContentInLayout(c.Layout, "Feedback 404", "", user, writer, "404.html", nil, 404)
		}
	}
}

func PrepareContext(r *Router, user *User, path, flash string, writer http.ResponseWriter, request *http.Request) *Context {
	c := Context{}
	c.Writer = writer
	c.Request = request
	c.flash = flash
	c.Method = request.Method
	c.router = r
	c.User = user
	c.path = path
	c.Db = r.Db
	c.Layout = "application_layout.html"
	return &c
}
