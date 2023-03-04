package router

import (
	"encoding/json"
	"html/template"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/persist"
	"github.com/jmoiron/sqlx"
)

type Router struct {
	Template *template.Template
	Site     *Site
	Db       *sqlx.DB
}

func NewRouter(path string) *Router {
	r := Router{}
	r.Db = persist.MysqlConnection()

	var site Site
	jsonString := files.ReadFile(path)
	json.Unmarshal([]byte(jsonString), &site)
	r.Site = &site

	r.Template = LoadTemplates()

	return &r
}
