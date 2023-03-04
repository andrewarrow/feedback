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
	Site     *Site
	Db       *sqlx.DB
	Location *time.Location
}

func NewRouter(path string) *Router {
	r := Router{}
	//r.Db = persist.MysqlConnection()
	r.Db = persist.PostgresConnection()

	r.Location, _ = time.LoadLocation("utc")
	time.Local = r.Location

	var site Site
	jsonString := persist.SchemaJson(r.Db)
	json.Unmarshal([]byte(jsonString), &site)
	r.Site = &site

	r.Template = LoadTemplates()

	return &r
}
