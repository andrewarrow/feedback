package sqlgen

import (
	"fmt"

	"github.com/andrewarrow/feedback/models"
)

func SqliteCreateTable(tableName string) string {
	sql := `CREATE TABLE %s (
  id INTEGER PRIMARY KEY,
	guid TEXT NOT NULL,
	created_at datetime CURRENT_TIMESTAMP,
	updated_at datetime CURRENT_TIMESTAMP
);`
	return fmt.Sprintf(sql, tableName)
}

func SqliteAlterTable(tableName string, model *models.Model) []string {
	sql := `ALTER TABLE %s ADD COLUMN %s %s default %s;`
	items := []string{}
	for _, field := range model.Fields {
		flavor, defaultString := SqlTypeAndDefault(field)
		a := fmt.Sprintf(sql, tableName, field.Name, flavor, defaultString)
		items = append(items, a)
	}
	return items
}

func SqlTypeAndDefault(f *models.Field) (string, string) {
	flavor := "TEXT"
	defaultString := "''"
	if f.Flavor == "int" {
		flavor = "INTEGER"
		defaultString = "0"
	} else if f.Flavor == "text" {
		flavor = "TEXT"
	} else if f.Flavor == "timestamp" {
		flavor = "timestamp"
		defaultString = "CURRENT_TIMESTAMP"
	}
	if f.Null == "yes" {
		defaultString = "NULL"
	}
	return flavor, defaultString
}
