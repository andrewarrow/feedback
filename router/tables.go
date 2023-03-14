package router

import (
	"fmt"
	"os"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/andrewarrow/feedback/util"
	"github.com/jmoiron/sqlx"
)

func MakeTables(db *sqlx.DB, models []*models.Model) {
	for _, model := range models {
		MakeTable(db, model)
	}
}

func MakeTable(db *sqlx.DB, model *models.Model) {
	prefix := os.Getenv("FEEDBACK_NAME")
	tableName := prefix + "_" + util.Plural(model.Name)
	//c.Db.Exec(sqlgen.MysqlCreateTable(tableName))
	db.Exec(sqlgen.PgCreateTable(tableName))
	sql := `ALTER TABLE %s ADD COLUMN %s %s default %s;`
	for _, field := range model.Fields {
		flavor, defaultString := field.SqlTypeAndDefault()
		db.Exec(fmt.Sprintf(sql, tableName, field.Name, flavor, defaultString))
		if field.Index == "yes" {
			sql := `create index %s_%s_index on %s(%s);`
			db.Exec(fmt.Sprintf(sql, tableName, field.Name, tableName, field.Name))
		} else if field.Index == "unique" {
			sql := `create unique index %s_%s_index on %s(%s);`
			db.Exec(fmt.Sprintf(sql, tableName, field.Name, tableName, field.Name))
		}
	}
}
