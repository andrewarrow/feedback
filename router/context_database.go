package router

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/models"
	"github.com/xeonx/timeago"
)

func (c *Context) FindModel(name string) *models.Model {
	return c.Router.FindModel(name)
}

func (r *Router) FindModel(name string) *models.Model {
	return r.Site.FindModel(name)
}

func CastFields(model *models.Model, m map[string]any) {
	if DB_FLAVOR == "pg" {
		CastFieldsPg(model, m)
	} else {
		CastFieldsSqlite(model, m)
	}
}
func CastFieldsPg(model *models.Model, m map[string]any) {
	cfg := timeago.English
	cfg.Max = 9223372036854775807

	if len(m) == 0 {
		return
	}
	for _, field := range model.Fields {
		if field.Flavor == "timestamp" && m[field.Name] != nil {
			tm := m[field.Name].(time.Time)
			ago := cfg.Format(tm)
			m[field.Name] = tm.Unix()
			m[field.Name+"_human"] = tm.Format(models.HUMAN)
			m[field.Name+"_date"] = tm.Format(models.HUMAN_DATE)
			m[field.Name+"_full_month"] = tm.Format(models.FULL_MONTH_DATE)
			m[field.Name+"_ago"] = ago
		} else if field.Flavor == "int" && m[field.Name] != nil {
			m[field.Name] = m[field.Name].(int64)
		} else if field.Flavor == "bigint" && m[field.Name] != nil {
			m[field.Name] = m[field.Name].(int64)
		} else if field.Flavor == "smallint" && m[field.Name] != nil {
			m[field.Name] = m[field.Name].(int16)
		} else if field.Flavor == "float" && m[field.Name] != nil {
			m[field.Name] = m[field.Name].(float64)
		} else if field.Flavor == "bool" && m[field.Name] != nil {
			m[field.Name] = m[field.Name].(bool)
		} else if field.Flavor == "double" && m[field.Name] != nil {
			m[field.Name] = m[field.Name].(float64)
		} else if field.Flavor == "numeric" && m[field.Name] != nil {
			floatString := string(m[field.Name].([]byte))
			result, _ := strconv.ParseFloat(floatString, 64)
			m[field.Name] = result
		} else if field.Flavor == "json" && m[field.Name] != nil {
			var temp map[string]any
			asString, ok := m[field.Name].(string)
			if ok {
				json.Unmarshal([]byte(asString), &temp)
			} else {
				json.Unmarshal(m[field.Name].([]byte), &temp)
			}
			m[field.Name] = temp
		} else if field.Flavor == "json_list" && m[field.Name] != nil {
			var temp []any
			asString, ok := m[field.Name].(string)
			if ok {
				json.Unmarshal([]byte(asString), &temp)
			} else {
				json.Unmarshal(m[field.Name].([]byte), &temp)
			}
			m[field.Name] = temp
		} else if field.Flavor == "list" {
			s := fmt.Sprintf("%s", m[field.Name])
			tokens := strings.Split(s, ",")
			m[field.Name] = tokens
		} else if m[field.Name] == nil {
			// to nothing, leave it nil
		} else {
			m[field.Name] = fmt.Sprintf("%s", m[field.Name])
		}
	}
}

var layout = "2006-01-02 15:04:05.999999-07:00"

func CastFieldsSqlite(model *models.Model, m map[string]any) {
	cfg := timeago.English
	cfg.Max = 9223372036854775807

	if len(m) == 0 {
		return
	}
	for _, field := range model.Fields {
		if field.Flavor == "timestamp" && m[field.Name] != nil {
			s, ok := m[field.Name].(string)
			tm := time.Now()
			if ok {
				tm, _ = time.Parse(layout, s)
			} else {
				s, _ := m[field.Name].(int64)
				tm = time.Unix(s, 0)
			}
			ago := cfg.Format(tm)
			m[field.Name] = tm.Unix()
			m[field.Name+"_human"] = tm.Format(models.HUMAN)
			m[field.Name+"_ago"] = ago
		} else if field.Flavor == "int" && m[field.Name] != nil {
			m[field.Name] = m[field.Name].(int64)
		} else if m[field.Name] == nil {
			// to nothing, leave it nil
		} else {
			m[field.Name] = fmt.Sprintf("%s", m[field.Name])
		}
	}
}
