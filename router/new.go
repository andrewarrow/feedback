package router

import (
	"encoding/json"
	"html/template"
	"sync"

	"github.com/andrewarrow/feedback/persist"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	Template       *template.Template
	Site           *FeedbackSite
	Db             *sqlx.DB
	Paths          map[string]func(*Context, string, string)
	BeforeCreate   map[string]func(*Context)
	AfterCreate    map[string]func(*Context, string)
	PathLock       sync.Mutex
	AfterLock      sync.Mutex
	BeforeLock     sync.Mutex
	DefaultLayout  string
	BearerAuthFunc func(*Context) map[string]any
	CookieAuthFunc func(*Context) map[string]any
}

func NewRouter(dbEnvVarName string, jsonBytes []byte) *Router {
	r := Router{}
	//r.Db = persist.MysqlConnection()
	r.Db = persist.PostgresConnection(dbEnvVarName)
	r.Paths = map[string]func(*Context, string, string){}
	r.BeforeCreate = map[string]func(*Context){}
	r.AfterCreate = map[string]func(*Context, string){}
	r.Paths["/"] = handleWelcome
	r.Paths["models"] = handleModels
	r.Paths["fields"] = handleFields
	r.Paths["sessions"] = handleSessions
	r.Paths["users"] = handleUsers
	r.Paths["about"] = handleAbout
	r.Paths["stats"] = handleStats
	r.Paths["admin"] = handleAdmin
	r.Paths["api"] = handleApi
	r.AfterCreate["user"] = afterCreateUser
	r.DefaultLayout = "application_layout.html"
	r.BearerAuthFunc = r.bearerAuth
	r.CookieAuthFunc = r.cookieAuth

	var site FeedbackSite
	if jsonBytes == nil {
		jsonBytes = []byte(persist.DefaultJsonModels())
	}
	json.Unmarshal(jsonBytes, &site)
	for _, m := range site.Models {
		m.EnsureIdAndCreatedAt()
	}
	r.Site = &site
	if r.Db != nil {
		go MakeTables(r.Db, r.Site.Models)
	}

	r.Template = LoadTemplates()

	return &r
}

func (r *Router) ToContext() *Context {
	c := Context{}
	//c.Writer = writer
	//c.Request = request
	c.Router = r
	c.Db = r.Db
	return &c
}

func (r *Router) pathFuncToRun(key string) func(*Context, string, string) {
	r.PathLock.Lock()
	defer r.PathLock.Unlock()
	return r.Paths[key]
}

func (r *Router) afterFuncToRun(key string) func(*Context, string) {
	r.AfterLock.Lock()
	defer r.AfterLock.Unlock()
	return r.AfterCreate[key]
}

func (r *Router) beforeFuncToRun(key string) func(*Context) {
	r.BeforeLock.Lock()
	defer r.BeforeLock.Unlock()
	return r.BeforeCreate[key]
}
