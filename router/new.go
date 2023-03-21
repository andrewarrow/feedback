package router

import (
	"encoding/json"
	"html/template"

	"github.com/andrewarrow/feedback/persist"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	Template    *template.Template
	Site        *FeedbackSite
	Db          *sqlx.DB
	Paths       map[string]func(*Context, string, string)
	AfterCreate map[string]func(*Context, string)
}

func NewRouter() *Router {
	r := Router{}
	//r.Db = persist.MysqlConnection()
	r.Db = persist.PostgresConnection()
	r.Paths = map[string]func(*Context, string, string){}
	r.AfterCreate = map[string]func(*Context, string){}
	r.Paths["/"] = handleWelcome
	r.Paths["models"] = handleModels
	r.Paths["sessions"] = handleSessions
	r.Paths["users"] = handleUsers
	r.Paths["about"] = handleAbout
	r.Paths["stats"] = handleStats
	r.Paths["tailwind"] = handleTailwind
	r.AfterCreate["user"] = afterCreateUser

	var site FeedbackSite
	jsonString := persist.SchemaJson(r.Db)
	json.Unmarshal([]byte(jsonString), &site)
	for _, m := range site.Models {
		m.EnsureIdAndCreatedAt()
	}
	r.Site = &site
	MakeTables(r.Db, r.Site.Models)

	r.Template = LoadTemplates()

	return &r
}
