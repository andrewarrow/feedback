package router

import (
	"fmt"
)

func (c *Context) MakeCells(list []any, headers []string, thing any, prefix string) [][]any {
	cells := [][]any{}
	for _, row := range list {
		thisRow := []any{}
		for j, _ := range headers {
			name := fmt.Sprintf("%s_col%d", prefix, j+1)
			templateVars := map[string]any{}
			templateVars["row"] = row
			if thing != nil {
				templateVars["params"] = thing
			}
			cell := c.Template(name, templateVars)
			thisRow = append(thisRow, cell)
		}
		cells = append(cells, thisRow)
	}
	return cells
}
