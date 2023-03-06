package router

import (
	"encoding/json"
	"html/template"

	"github.com/andrewarrow/feedback/persist"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	Template *template.Template
	Site     *FeedbackSite
	Db       *sqlx.DB
}

func NewRouter() *Router {
	r := Router{}
	//r.Db = persist.MysqlConnection()
	r.Db = persist.PostgresConnection()

	var site FeedbackSite
	jsonString := persist.SchemaJson(r.Db)
	json.Unmarshal([]byte(jsonString), &site)
	r.Site = &site
	MakeTables(r.Db, r.Site.Models)

	r.Template = LoadTemplates()

	return &r
}
