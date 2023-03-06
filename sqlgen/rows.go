package sqlgen

import (
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/util"
)

func UpdateRow(model *models.Model) string {
	tableName := util.Plural(model.Name)
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

func InsertRow(tableName string, fields []*models.Field) (string, []any) {
	buffer := []string{"INSERT INTO "}
	buffer = append(buffer, tableName+" (")

	cols := []string{}
	for _, field := range fields {
		cols = append(cols, field.Name)
	}
	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, ") values (")
	cols = []string{}
	params := []any{}
	for i, field := range fields {
		cols = append(cols, fmt.Sprintf("$%d", i+1))
		params = append(params, field.RandomValue())
	}
	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, ")")

	return strings.Join(buffer, ""), params
}
