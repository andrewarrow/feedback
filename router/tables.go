package router

import (
	"encoding/json"
	"fmt"
	"strings"

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
	tableName := model.TableName()
	//c.Db.Exec(sqlgen.MysqlCreateTable(tableName))
	//_, msg := db.Exec(sqlgen.PgCreateTable(tableName))
	db.Exec(sqlgen.PgCreateTable(tableName))

	sql := `ALTER TABLE %s ADD COLUMN %s %s default %s;`
	for _, field := range model.Fields {
		flavor, defaultString := field.SqlTypeAndDefault()
		//_, msg := db.Exec(fmt.Sprintf(sql, tableName, field.Name, flavor, defaultString))
		db.Exec(fmt.Sprintf(sql, tableName, field.Name, flavor, defaultString))
		//fmt.Println(msg, flavor)
		if field.Index == "yes" {
			sql := `create index CONCURRENTLY %s_%s_index on %s(%s);`
			db.Exec(fmt.Sprintf(sql, tableName, field.Name, tableName, field.Name))
		} else if field.Index == "unique" {
			sql := `create unique index %s_%s_index on %s(%s);`
			db.Exec(fmt.Sprintf(sql, tableName, field.Name, tableName, field.Name))
		} else if strings.HasPrefix(field.Index, "unique_two") {
			tokens := strings.Split(field.Index, ":")
			fields := strings.Split(tokens[1], ",")
			field1 := fields[0]
			field2 := fields[1]
			sql := `create unique index %s_%s_%s_index on %s(%s,%s);`
			db.Exec(fmt.Sprintf(sql, tableName, field1, field2, tableName, field1, field2))
		}
	}
}

func ModelsToBytes(list []*models.Model) []byte {
	site := FeedbackSite{}
	site.Models = list
	asBytes, _ := json.Marshal(site)
	return asBytes
}

func MakeGuidsInTables(db *sqlx.DB, models []*models.Model) {
	for _, model := range models {
		MakeGuidsInTable(db, model)
	}
}

func MakeGuidsInTable(db *sqlx.DB, model *models.Model) {
	tableName := model.TableName()
	sql := `update %s set guid=$1 where id=$2;`
	for i := 1; i < 1000; i++ {
		guid := util.PseudoUuid()
		s := fmt.Sprintf(sql, tableName)
		fmt.Println(s, guid, i)
		db.Exec(s, guid, i)
	}
}
