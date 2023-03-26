package sqlgen

import (
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/models"
)

func UpdateRow(model *models.Model) string {
	tableName := model.TableName()
	buffer := []string{"UPDATE "}
	buffer = append(buffer, tableName+" set ")

	cols := []string{}
	for i, field := range model.Fields {
		cols = append(cols, fmt.Sprintf("%s=$%d", field.Name, i+1))
	}
	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, fmt.Sprintf(" where guid=$%d", len(model.Fields)+1))

	return strings.Join(buffer, "")
}

func InsertRowNoRandomDefaults(tableName string,
	fields []*models.Field,
	override map[string]any) (string, []any) {
	return insertRow(false, tableName, fields, override)
}

func InsertRowWithRandomDefaults(tableName string,
	fields []*models.Field,
	override map[string]any) (string, []any) {
	return insertRow(true, tableName, fields, override)
}

func insertRow(random bool, tableName string,
	fields []*models.Field,
	override map[string]any) (string, []any) {

	buffer := []string{"INSERT INTO "}
	buffer = append(buffer, tableName+" (")

	cols := []string{}
	for _, field := range fields {
		if field.Name == "id" || field.Name == "created_at" {
			continue
		}
		cols = append(cols, field.Name)
	}
	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, ") values (")
	cols = []string{}
	params := []any{}
	count := 1
	for _, field := range fields {
		if field.Name == "id" || field.Name == "created_at" {
			continue
		}
		cols = append(cols, fmt.Sprintf("$%d", count))
		count++
		val := override[field.Name]
		if val == nil {
			val = ""
			if random {
				val = field.RandomValue()
			}
		}
		params = append(params, val)
	}
	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, ")")

	return strings.Join(buffer, ""), params
}
