package router

import (
	"encoding/json"
	"fmt"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/jmoiron/sqlx"
)

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

func ModelsToBytes(list []*models.Model) []byte {
	site := FeedbackSite{}
	site.Models = list
	asBytes, _ := json.Marshal(site)
	return asBytes
}
