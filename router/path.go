package router

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/andrewarrow/feedback/util"
)

func (r *Router) PlaceContentInLayoutMap(layoutMap map[string]any, flash string, user map[string]any, filename string, vars any) {
	content := new(bytes.Buffer)
	r.Template.ExecuteTemplate(content, filename, vars)

	layoutMap["footer"] = r.Site.Footer
	layoutMap["site_name"] = r.Site.Title
	layoutMap["flash"] = flash
	layoutMap["user"] = user
	layoutMap["dev_mode"] = os.Getenv("DEV_MODE") != ""
	layoutMap["content"] = template.HTML(content.String())
}

func (r *Router) SendContentInLayout(doZip bool, layout string, layoutMap map[string]any, flash string, user map[string]any, writer http.ResponseWriter,
	filename string, contentVars any, status int) {
	r.PlaceContentInLayoutMap(layoutMap, flash, user, filename, contentVars)
	r.sendZippy(doZip, layout, layoutMap, writer, status)
}

func (r *Router) cookieAuth(c *Context) map[string]any {
	cookie, err := c.Request.Cookie("user")
	var user map[string]any
	if err == nil && cookie.Value != "" {
		user = r.LookupUserByToken(cookie.Value)
	}
	return user
}

func (r *Router) bearerAuth(c *Context) map[string]any {
	var user map[string]any
	auth := util.GetHeader("Authorization", c.Request)
	if auth != "" {
		tokens := strings.Split(auth, " ")
		if len(tokens) == 2 {
			guid := tokens[1]
			user = r.LookupUserByToken(guid)
		}
	}
	return user
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	c := PrepareContext(r, nil, "/", "", writer, request)

	var user map[string]any
	if r.CookieAuthFunc != nil {
		user = r.CookieAuthFunc(c)
	}
	if user == nil && r.BearerAuthFunc != nil {
		user = r.BearerAuthFunc(c)
	}

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
		r.HandleAsset("/assets/other/robots.txt", writer, request)
	} else if strings.HasPrefix(path, "/favicon.ico") {
		r.HandleAsset("/assets/other/favicon.ico", writer, request)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer, request)
	} else if strings.HasSuffix(path, "/") {
		http.Redirect(writer, request, path[0:len(path)-1], 301)
	} else {
		fmt.Println(request.Method, path)
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
			c.LayoutMap["title"] = "404 not found"
			r.SendContentInLayout(false, c.Layout, c.LayoutMap, "", user, writer, "404.html", nil, 404)
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
	c.LayoutMap = map[string]any{}
	return &c
}
