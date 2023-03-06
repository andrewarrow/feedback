package sqlgen

import (
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/models"
)

func InsertRow(tableName string, fields []models.Field) (string, []any) {
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
