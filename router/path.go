package router

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/util"
)

type LayoutVars struct {
	Title    string
	SiteName string
	User     map[string]any
	Footer   string
	Content  template.HTML
	Flash    string
}

func (r *Router) PlaceContentInLayoutVars(title, flash string, user map[string]any, filename string, vars any) *LayoutVars {
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

func (r *Router) SendContentInLayout(layout, title, flash string, user map[string]any, writer http.ResponseWriter,
	filename string, contentVars any, status int) {
	vars := r.PlaceContentInLayoutVars(title, flash, user, filename, contentVars)
	writer.WriteHeader(status)
	r.Template.ExecuteTemplate(writer, layout, vars)
}

func (r *Router) cookieAuth(request *http.Request) map[string]any {
	cookie, err := request.Cookie("user")
	var user map[string]any
	if err == nil && cookie.Value != "" {
		user = r.LookupUser(cookie.Value)
	}
	return user
}

func (r *Router) bearerAuth(request *http.Request) map[string]any {
	var user map[string]any
	auth := util.GetHeader("Authorization", request)
	if auth != "" {
		tokens := strings.Split(auth, " ")
		if len(tokens) == 2 {
			guid := tokens[1]
			user = r.LookupUser(guid)
		}
	}
	return user
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	var user map[string]any
	if r.CookieAuthFunc != nil {
		user = r.CookieAuthFunc(request)
	}
	user = r.BearerAuthFunc(request)

	cookie, err := request.Cookie("flash")
	flash := ""
	if err == nil && cookie.Value != "" {
		flash = cookie.Value
		removeFlash(writer)
	}

	if path == "/" {
		c := PrepareContext(r, user, "/", flash, writer, request)
		r.pathFuncToRun("/")(c, "", "")
	} else if strings.HasPrefix(path, "/robots.txt") {
		r.HandleAsset("/assets/robots.txt", writer, request)
	} else if strings.HasPrefix(path, "/favicon.ico") {
		r.HandleAsset("/assets/favicon.ico", writer, request)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer, request)
	} else if strings.HasSuffix(path, "/") {
		http.Redirect(writer, request, path[0:len(path)-1], 301)
	} else {
		path = path + "/"
		c := PrepareContext(r, user, path, flash, writer, request)
		c.tokens = strings.Split(path, "/")
		contentType := util.GetHeader("Content-Type", request)
		if c.Method == "POST" && contentType == "application/x-www-form-urlencoded" {
			c.ReadFormPost()
		}
		handleContext(c)
		if c.UserRequired && len(c.User) == 0 {
			http.Redirect(c.Writer, c.Request, "/sessions/new/", 302)
			return
		}
		if c.NotFound && c.Layout != "json" {
			r.SendContentInLayout(c.Layout, "Feedback 404", "", user, writer, "404.html", nil, 404)
		} else if c.NotFound && c.Layout == "json" {
			c.SendContentAsJsonMessage("not found", 404)
		}
	}
}

func PrepareContext(r *Router, user map[string]any, path, flash string, writer http.ResponseWriter, request *http.Request) *Context {
	c := Context{}
	c.Writer = writer
	c.Request = request
	c.flash = flash
	c.Method = request.Method
	c.router = r
	c.User = user
	c.path = path
	c.Db = r.Db
	c.Layout = r.DefaultLayout
	return &c
}
