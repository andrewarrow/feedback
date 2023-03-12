package router

import (
	"encoding/json"
	"html/template"

	"github.com/andrewarrow/feedback/persist"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	Template          *template.Template
	Site              *FeedbackSite
	Db                *sqlx.DB
	Paths             map[string]func(*Context, string, string)
	UserRequiredPaths map[string]*UserRequired
}

func NewRouter() *Router {
	r := Router{}
	//r.Db = persist.MysqlConnection()
	r.Db = persist.PostgresConnection()
	r.Paths = map[string]func(*Context, string, string){}
	r.Paths["models"] = handleModels
	r.Paths["sessions"] = handleSessions
	r.Paths["users"] = handleUsers
	r.Paths["about"] = handleAbout

	r.UserRequiredPaths = map[string]*UserRequired{}
	r.UserRequiredPaths["/sessions/new/"] = NewUserRequired("GET", "==")
	r.UserRequiredPaths["/sessions/"] = NewUserRequired("POST", "==")
	r.UserRequiredPaths["/about/"] = NewUserRequired("GET", "==")
	r.UserRequiredPaths["/users/"] = NewUserRequired("*", "prefix")

	var site FeedbackSite
	jsonString := persist.SchemaJson(r.Db)
	json.Unmarshal([]byte(jsonString), &site)
	r.Site = &site
	MakeTables(r.Db, r.Site.Models)

	r.Template = LoadTemplates()

	return &r
}
