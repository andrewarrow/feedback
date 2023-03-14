package router

import (
	"fmt"
	"os"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/jmoiron/sqlx"
)

func TableName(name string) string {
	prefix := os.Getenv("FEEDBACK_NAME")
	return prefix + "_" + name
}

func MakeTables(db *sqlx.DB, models []*models.Model) {
	for _, model := range models {
		MakeTable(db, model)
	}
}

func MakeTable(db *sqlx.DB, model *models.Model) {
	tableName := model.TableName()
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
