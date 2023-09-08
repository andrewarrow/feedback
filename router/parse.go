package router

import (
	"strconv"
)

func ParseNumbers(c *Context, cols []string, editable map[string]string) {
	for _, item := range cols {
		if editable[item] == "int" {
			c.Params[item], _ = strconv.Atoi(c.Params[item].(string))
		} else if editable[item] == "float" {
			c.Params[item], _ = strconv.ParseFloat(c.Params[item].(string), 64)
		}
	}
}

func IsEditable(item string, editable map[string]string) bool {
	if editable[item] != "string" &&
		editable[item] != "text" &&
		editable[item] != "int" &&
		editable[item] != "float" &&
		editable[item] != "json" &&
		editable[item] != "select" &&
		editable[item] != "select-multi" &&
		editable[item] != "bool" {
		return false
	}
	return true
}
