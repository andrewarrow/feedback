package router

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/andrewarrow/feedback/util"
)

func (c *Context) SendContentAsJson(thing any, status int) {
	asBytes, _ := json.Marshal(thing)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Cache-Control", "none")
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(asBytes)))
	c.Writer.WriteHeader(status)
	c.Writer.Write(asBytes)
}

func (c *Context) SendContentAsJsonMessage(message string, status int) {
	m := map[string]any{"info": message}
	c.SendContentAsJson(m, status)
}

func (c *Context) CreateRowFromJson(modelString string) string {
	model := c.FindModel(modelString)
	tableName := model.TableName()
	sql, params := sqlgen.InsertRowNoRandomDefaults(tableName, model.Fields, c.Params)
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (c *Context) ValidateJsonForModel(modelString string) string {
	c.Params = c.ReadBodyIntoJson()
	model := c.FindModel(modelString)

	for _, field := range model.Fields {
		if field.Required == "yes" {
			if c.Params[field.Name] == nil {
				return "missing " + field.Name
			}
		} else if strings.HasPrefix(field.Required, "if") {
			tokens := strings.Split(field.Required, " ")
			value := tokens[1]
			if strings.HasPrefix(value, "!") {
				value = value[1:]
				if c.Params[value] == nil && c.Params[field.Name] == nil {
					return "missing " + value + " or " + field.Name
				}
			}
		}
	}

	for _, field := range model.Fields {
		if field.Regex == "" {
			continue
		}
		if c.Params[field.Name] == nil {
			continue
		}

		val := c.Params[field.Name].(string)

		re := regexp.MustCompile(field.Regex)
		if !re.MatchString(val) {
			return "wrong format " + field.Name
		}
	}

	guid := util.PseudoUuid()
	c.Params["guid"] = guid

	return ""
}
