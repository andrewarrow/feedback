package router

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/xeonx/timeago"
)

type WelcomeVars struct {
	Rows []map[string]any
}

const layout = "2006-01-02 15:04:05"

func WelcomeIndexVars(db *sqlx.DB, location *time.Location) *WelcomeVars {
	vars := WelcomeVars{}
	vars.Rows = []map[string]any{}
	rows, _ := db.Queryx("SELECT * FROM stories ORDER BY created_at desc limit 30")
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		for k, v := range m {
			m[k] = fmt.Sprintf("%s", v)
		}
		tokens := strings.Split(m["created_at"].(string), ".")
		timestamp := tokens[0]
		tm, _ := time.Parse(layout, timestamp)
		tm = tm.Add(time.Hour * 8)
		m["timestamp"] = timestamp
		m["ago"] = timeago.English.Format(tm)
		if m["url"].(string) == "" {
			m["url"] = "/stories/" + m["guid"].(string) + "/"
		}
		vars.Rows = append(vars.Rows, m)
	}
	return &vars
}
