package router

import (
	"fmt"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/andrewarrow/feedback/util"
	"github.com/jmoiron/sqlx"
)

func (r *Router) ResetDatabase() {
	for _, model := range r.Site.Models {
		tableName := util.Plural(model.Name)
		r.Db.Exec("drop table " + tableName)
	}
	r.Db.Exec("drop table feedback_schema")
	fmt.Println("done.")
}

func MakeTables(db *sqlx.DB, models []*models.Model) {
	for _, model := range models {
		MakeTable(db, model)
	}
}

func MakeTable(db *sqlx.DB, model *models.Model) {
	tableName := util.Plural(model.Name)
	//c.db.Exec(sqlgen.MysqlCreateTable(tableName))
	db.Exec(sqlgen.PgCreateTable(tableName))
	flavor := "varchar(255)"
	defaultString := "''"
	sql := `ALTER TABLE %s ADD COLUMN %s %s default %s;`
	for _, field := range model.Fields {
		if field.Flavor == "text" {
			flavor = "text"
			defaultString = "''"
		} else if field.Flavor == "int" {
			flavor = "int"
			defaultString = "0"
		} else if field.Flavor == "uuid" {
			flavor = "varchar(255)"
			defaultString = "''"
		}
		db.Exec(fmt.Sprintf(sql, tableName, field.Name, flavor, defaultString))
		if field.Index == "yes" {
			sql := `create index %s_index on %s(%s);`
			db.Exec(fmt.Sprintf(sql, field.Name, tableName, field.Name))
		} else if field.Index == "unique" {
			sql := `create unique index %s_index on %s(%s);`
			db.Exec(fmt.Sprintf(sql, field.Name, tableName, field.Name))
		}
	}
}
