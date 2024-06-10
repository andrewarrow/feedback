package sqlgen

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/models"
	"github.com/lib/pq"
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

func InsertRowNoRandomDefaults(dbFlavor, tableName string,
	fields []*models.Field,
	override map[string]any) (string, []any) {
	return insertRow(false, dbFlavor, tableName, fields, override)
}

func InsertRowWithRandomDefaults(dbFlavor, tableName string,
	fields []*models.Field,
	override map[string]any) (string, []any) {
	return insertRow(true, dbFlavor, tableName, fields, override)
}

func insertRow(random bool, dbFlavor, tableName string,
	fields []*models.Field,
	override map[string]any) (string, []any) {

	now := time.Now()
	override["created_at"] = now
	override["updated_at"] = now

	buffer := []string{"INSERT INTO "}
	buffer = append(buffer, tableName+" (")

	cols := []string{}
	for _, field := range fields {
		if field.Name == "id" {
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
		if field.Name == "id" {
			continue
		}
		cols = append(cols, fmt.Sprintf("$%d", count))
		count++
		val := override[field.Name]
		isNullString, _ := val.(string)
		if val == nil {
			if random {
				val = field.RandomValue()
			} else {
				val = field.SaneDefault()
			}
		}
		if field.Flavor == "list" && val != nil {
			list := []string{}
			thing1, isArrayAny := val.([]any)
			thing2, isArrayString := val.([]string)
			thing3, isString := val.(string)

			if isString {
				list = append(list, strings.ToLower(thing3))
			} else if isArrayAny {
				for _, s := range thing1 {
					list = append(list, strings.ToLower(s.(string)))
				}
			} else if isArrayString {
				for _, s := range thing2 {
					list = append(list, strings.ToLower(s))
				}
			}
			val = strings.Join(list, ",")
		} else if field.Flavor == "json" {
			asBytes, _ := json.Marshal(val)
			val = string(asBytes)
		} else if field.Flavor == "array" {
			val = pq.Array([0]string{})
		} else if field.Flavor == "timestamp" {
			ts, ok := val.(time.Time)
			if ok && dbFlavor == "sqlite" {
				val = ts.Unix()
			}
		} else if field.Flavor == "json_list" {
			asBytes, _ := json.Marshal(val)
			val = string(asBytes)
		} else if isNullString == "null" {
			var sqlNullString sql.NullString
			val = sqlNullString
		}
		params = append(params, val)
	}
	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, ")")

	sql := strings.Join(buffer, "")
	if os.Getenv("DEBUG") == "1" {
		fmt.Println(sql, params)
	}
	return sql, params
}

func UpdateRowFromParams(dbFlavor, tableName string,
	fields []*models.Field,
	override map[string]any, where string) (string, []any) {

	params := []any{}
	buffer := []string{"UPDATE "}
	buffer = append(buffer, tableName+" set ")

	cols := []string{}
	count := 1
	for _, field := range fields {
		if field.Name == "id" || field.Name == "created_at" || field.Name == "updated_at" {
			continue
		}
		cols = append(cols, fmt.Sprintf("%s=$%d", field.Name, count))
		count++
		val := override[field.Name]
		isNullString, _ := val.(string)
		if field.Flavor == "list" {
			list := fixListItems(val)
			val = strings.Join(list, ",")
		} else if field.Flavor == "json" {
			asBytes, _ := json.Marshal(val)
			val = string(asBytes)
		} else if field.Flavor == "array" {
			val = pq.Array([0]string{})
		} else if field.Flavor == "timestamp" {
			ts, ok := val.(time.Time)
			if ok && dbFlavor == "sqlite" {
				val = ts.Unix()
			}
		} else if field.Flavor == "json_list" {
			asBytes, _ := json.Marshal(val)
			val = string(asBytes)
		} else if isNullString == "null" {
			var sqlNullString sql.NullString
			val = sqlNullString
		}
		params = append(params, val)
	}
	cols = append(cols, fmt.Sprintf("updated_at=$%d", count))
	params = append(params, time.Now())

	buffer = append(buffer, strings.Join(cols, ","))
	buffer = append(buffer, fmt.Sprintf(" %s$%d", where, count+1))
	return strings.Join(buffer, ""), params
}

func fixListItems(val any) []string {
	s, ok := val.(string)
	if ok {
		return []string{s}
	}
	list := []string{}
	items, ok := val.([]any)
	if ok {
		for _, s := range items {
			list = append(list, strings.ToLower(s.(string)))
		}
		return list
	}
	return val.([]string)
}
