package sqlgen

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/util"
	"github.com/brianvoe/gofakeit/v6"
)

func InsertRow(tableName string, fields []models.Field) string {
	buffer := []string{"INSERT INTO "}
	buffer = append(buffer, tableName+" (")

	cols := []string{}
	for _, field := range fields {
		cols = append(cols, field.Name)
	}
	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, ") values (")

	cols = []string{}
	for _, field := range fields {
		var val string
		if field.Flavor == "uuid" {
			val = "'" + util.PseudoUuid() + "'"
		} else if field.Flavor == "username" {
			val = "'" + gofakeit.Username() + "'"
		} else if field.Flavor == "int" {
			val = fmt.Sprintf("%d", rand.Intn(999))
		} else {
			val = "'" + gofakeit.Word() + "'"
		}
		cols = append(cols, val)
	}
	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, ");")

	return strings.Join(buffer, "")
}
