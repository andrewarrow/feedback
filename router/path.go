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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r, filename)
		}
	}()
	content := new(bytes.Buffer)

	contentVars, ok := vars.(map[string]any)
	if ok {
		contentVars["user"] = user
	}

	t := r.GetLiveOrCachedTemplate(filename)
	t.Execute(content, vars)

	layoutMap["content_vars"] = vars
	layoutMap["footer"] = r.Site.Footer
	layoutMap["site_name"] = r.Site.Title
	layoutMap["flash"] = flash
	layoutMap["user"] = user
	layoutMap["viewport"] = viewport
	layoutMap["wasm"] = MakeWasmScript(BuildTag, filename)
	layoutMap["USE_LIVE_TEMPLATES"] = UseLiveTemplates
	layoutMap["dev_mode"] = os.Getenv("DEV_MODE") != ""
	layoutMap["content"] = template.HTML(content.String())
}

func (r *Router) SendContentInLayout(doZip bool, layout string, layoutMap map[string]any, flash string, user map[string]any, writer http.ResponseWriter,
	filename string, contentVars any, status int) {
	r.PlaceContentInLayoutMap(layoutMap, flash, user, filename, contentVars)
	r.sendZippy(doZip, layout, layoutMap, writer, status)
}

func (r *Router) cookieAuth(c *Context) map[string]any {
	cookie, err := c.Request.Cookie("user_v2")
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
	r.BeforeAll(c)

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

	//CORS_DOMAIN := os.Getenv("CORS_DOMAIN")

	if path == "/" || path == "/"+r.Prefix || path == r.Prefix {
		c := PrepareContext(r, user, "/", flash, writer, request)
		r.pathFuncToRun("/")(c, "", "")
	} else if strings.HasPrefix(path, "/robots.txt") {
		r.HandleAsset("/assets/other/robots.txt", writer, request)
	} else if strings.HasPrefix(path, "/manifest.json") {
		r.HandleAsset("/assets/other/manifest.json", writer, request)
	} else if strings.HasPrefix(path, "/google") {
		r.HandleAsset("/assets/other"+path, writer, request)
	} else if strings.HasPrefix(path, "/favicon.ico") {
		r.HandleAsset("/assets/other/favicon.ico", writer, request)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer, request)
	} else if strings.HasPrefix(path, "/bucket") {
		r.HandleBucketAsset(path, writer, request)
	} else if strings.HasPrefix(path, "/"+r.Prefix+"/assets") {
		c.tokens = strings.Split(path, "/")
		newPath := "/" + strings.Join(c.tokens[2:], "/")
		fmt.Println(newPath)
		r.HandleAsset(newPath, writer, request)
	} else if strings.HasSuffix(path, "/") {
		http.Redirect(writer, request, path[0:len(path)-1], 301)
	} else if c.Method == "OPTIONS" {
		writer.Header().Set("Allow", "GET,POST,PUT,PATCH,DELETE")
		//writer.Header().Set("Access-Control-Allow-Origin", "http://*."+CORS_DOMAIN)
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' http://localhost")
		writer.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubdomains")

	} else {
		//writer.Header().Set("Access-Control-Allow-Origin", "http://*."+CORS_DOMAIN)
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		fmt.Println(request.Method, path)
		path = path + "/"
		if r.Prefix != "" {
			c.tokens = strings.Split(path, "/")
			path = "/" + strings.Join(c.tokens[2:], "/")
		}
		c := PrepareContext(r, user, path, flash, writer, request)
		c.tokens = strings.Split(path, "/")
		contentType := util.GetHeader("Content-Type", request)
		if c.Method == "POST" && contentType == "application/x-www-form-urlencoded" {
			c.ReadFormPost()
		}
		handleContext(c)
		if c.UserRequired && len(c.User) == 0 {
			returnPath := "/sessions/new"
			if r.Prefix != "" {
				returnPath = "/" + r.Prefix + "/sessions/new"
			}
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}
		r.NotFoundFunc(r, c)
	}
}

func Default404(r *Router, c *Context) {
	if c.NotFound && c.Layout != "json" {
		c.LayoutMap["title"] = "404 not found"
		r.SendContentInLayout(false, c.Layout, c.LayoutMap, "", c.User, c.Writer, "404.html", nil, 404)
	} else if c.NotFound && c.Layout == "json" {
		c.SendContentAsJsonMessage("not found", 404)
	}
}

func PrepareContext(r *Router, user map[string]any, path, flash string, writer http.ResponseWriter, request *http.Request) *Context {
	c := Context{}
	c.Writer = writer
	c.Request = request
	c.flash = flash
	c.Method = request.Method
	c.Router = r
	c.User = user
	c.path = path
	c.Db = r.Db
	c.Layout = r.DefaultLayout
	c.LayoutMap = map[string]any{}
	return &c
}
