package router

import (
	"encoding/json"
	"html/template"
	"time"

	_ "time/tzdata"

	"github.com/andrewarrow/feedback/persist"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	Template *template.Template
	Site     *FeedbackSite
	Db       *sqlx.DB
	Location *time.Location
}

func NewRouter() *Router {
	r := Router{}
	//r.Db = persist.MysqlConnection()
	r.Db = persist.PostgresConnection()

	r.Location, _ = time.LoadLocation("utc")
	time.Local = r.Location

	var site FeedbackSite
	jsonString := persist.SchemaJson(r.Db)
	json.Unmarshal([]byte(jsonString), &site)
	r.Site = &site
	MakeTables(r.Db, r.Site.Models)

	r.Template = LoadTemplates()

	return &r
}
