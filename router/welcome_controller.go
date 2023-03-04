package router

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type WelcomeVars struct {
	Rows []map[string]any
}

func WelcomeIndexVars(db *sqlx.DB) *WelcomeVars {
	vars := WelcomeVars{}
	vars.Rows = []map[string]any{}
	rows, _ := db.Queryx("SELECT * FROM stories ORDER BY created_at desc limit 30")
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		for k, v := range m {
			m[k] = fmt.Sprintf("%s", v)
		}
		vars.Rows = append(vars.Rows, m)
	}
	return &vars
}
