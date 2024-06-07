package router

import (
	"encoding/json"
	"html/template"
	"sync"

	"github.com/andrewarrow/feedback/persist"
	"github.com/jmoiron/sqlx"
)

var CustomFuncMap *template.FuncMap
var UseJsonToMakeTablesAndIndexes = true
var DB_FLAVOR = "pg"

type Router struct {
	Template        *template.Template
	Site            *FeedbackSite
	Db              *sqlx.DB
	WrangleDb       *sqlx.DB
	Prefix          string
	BucketPath      string
	Paths           map[string]func(*Context, string, string)
	BeforeCreate    map[string]func(*Context)
	AfterCreate     map[string]func(*Context, string)
	PathLock        sync.Mutex
	AfterLock       sync.Mutex
	BeforeLock      sync.Mutex
	DefaultLayout   string
	BearerAuthFunc  func(*Context) map[string]any
	CookieAuthFunc  func(*Context) map[string]any
	NotFoundFunc    func(*Router, *Context)
	BeforeAll       func(*Context)
	NotLoggedInPath string
}

func NewRouter(dbEnvVarName string, jsonBytes []byte) *Router {
	r := Router{}
	//r.Db = persist.MysqlConnection()
	if DB_FLAVOR == "pg" {
		r.Db = persist.PostgresConnection(dbEnvVarName)
		r.WrangleDb = persist.PostgresConnection("WRANGLE_DATABASE_URL")
	} else {
		r.Db = persist.SqliteConnection()
	}
	r.Paths = map[string]func(*Context, string, string){}
	r.BeforeCreate = map[string]func(*Context){}
	r.AfterCreate = map[string]func(*Context, string){}
	r.Paths["/"] = handleWelcome
	r.Paths["models"] = handleModels
	r.Paths["fields"] = handleFields
	r.Paths["sessions"] = HandleSessions
	r.Paths["users"] = handleUsers
	r.Paths["about"] = handleAbout
	r.Paths["stats"] = handleStats
	//r.Paths["admin"] = handleAdmin
	r.Paths["api"] = handleApi
	r.Paths["google"] = handleGoogle
	r.AfterCreate["user"] = afterCreateUser
	r.DefaultLayout = "application_layout.html"
	r.BeforeAll = r.beforeAllFunc
	r.BearerAuthFunc = r.bearerAuth
	r.CookieAuthFunc = r.cookieAuth
	r.NotFoundFunc = Default404

	var site FeedbackSite
	if jsonBytes == nil {
		jsonBytes = []byte(persist.DefaultJsonModels())
	}
	json.Unmarshal(jsonBytes, &site)
	for _, m := range site.Models {
		m.EnsureIdAndCreatedAt()
	}
	r.Site = &site
	if r.Db != nil && UseJsonToMakeTablesAndIndexes {
		go MakeTables(r.Db, r.Site.Models)
	}

	if CustomFuncMap == nil {
		tf := TemplateFunctions()
		CustomFuncMap = &tf
	}
	r.Template = LoadTemplates(*CustomFuncMap)

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
