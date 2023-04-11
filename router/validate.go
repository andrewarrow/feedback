package router

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/util"
)

func (c *Context) ValidateOneField(modelString, fieldString, value string) bool {
	model := c.FindModel(modelString)
	field := models.FindField(model, fieldString)
	if field.Regex == "" {
		return true
	}
	re := regexp.MustCompile(field.Regex)
	if !re.MatchString(value) {
		return false
	}
	return true
}

func (c *Context) ValidateCreate(modelString string) string {
	model := c.FindModel(modelString)
	return c.Validate(true, model.Fields)
}

func (c *Context) ValidateUpdate(modelString string) string {
	model := c.FindModel(modelString)
	list := []*models.Field{}
	for _, field := range model.Fields {
		if c.Params[field.Name] == nil {
			continue
		}
		list = append(list, field)
	}
	return c.Validate(false, list)
}

func (c *Context) Validate(create bool, fields []*models.Field) string {

	for _, field := range fields {
		if field.Flavor != "timestamp" {
			continue
		}
		if c.Params[field.Name] != nil {
			var t time.Time
			stringTime, ok := c.Params[field.Name].(string)
			if ok {
				intTime, _ := strconv.ParseInt(stringTime, 10, 64)
				t = time.Unix(intTime, 0)
			} else {
				t = time.Unix(int64(c.Params[field.Name].(float64)), 0)
			}
			c.Params[field.Name] = t
		}
	}

	for _, field := range fields {
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

	for _, field := range fields {
		if field.Regex == "" {
			continue
		}
		if field.Null == "yes" && c.Params[field.Name] == nil {
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

	if create {
		guid := util.PseudoUuid()
		c.Params["guid"] = guid
	}

	return ""
}
