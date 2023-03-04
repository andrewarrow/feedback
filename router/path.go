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
		c.method = request.Method
		c.router = r
		c.user = user
		c.path = path
		c.db = r.Db
		c.tokens = strings.Split(path, "/")
		c.userRequired = r.IsUserRequired(path, c.method)
		if c.userRequired && c.user == nil {
			http.Redirect(c.writer, c.request, "/sessions/new/", 302)
			return
		}
		handleContext(&c)
		if c.notFound {
			r.SendContentInLayout(user, writer, "404.html", nil, 404)
		}
	}
}
