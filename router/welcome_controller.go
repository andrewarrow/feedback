package router

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type WelcomeVars struct {
	Rows []*Story
}

func WelcomeIndexVars(db *sqlx.DB, order, domain string) *WelcomeVars {
	vars := WelcomeVars{}
	vars.Rows = []*Story{}
	params := []any{}
	sql := fmt.Sprintf("SELECT * FROM stories ORDER BY %s limit 30", order)
	if domain != "" {
		sql = fmt.Sprintf("SELECT * FROM stories where domain=$1 ORDER BY %s limit 30", order)
		params = append(params, domain)
	}
	rows, err := db.Queryx(sql, params...)
	if err != nil {
		return &vars
	}
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		story := storyFromMap(m)
		vars.Rows = append(vars.Rows, story)
	}
	return &vars
}
