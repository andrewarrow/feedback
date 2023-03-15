package router

import (
	"time"

	"github.com/andrewarrow/feedback/models"
	"github.com/xeonx/timeago"
)

func FixTime(model *models.Model, m *map[string]any) {
	for _, field := range model.Fields {
		if field.Flavor != "timestamp" {
			continue
		}
		tm := (*m)[field.Name].(time.Time)
		ago := timeago.English.Format(tm)
		(*m)[field.Name] = tm.Format(time.RFC1123)
		(*m)[field.Name+"_ago"] = ago
	}
}

func FixTimes(model *models.Model, rows []*map[string]any) {
	for _, row := range rows {
		FixTime(model, row)
	}
}
