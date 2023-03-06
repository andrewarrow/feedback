package router

import "github.com/jmoiron/sqlx"

type WelcomeVars struct {
	Rows []*Story
}

func WelcomeIndexVars(db *sqlx.DB, order, domain string) *WelcomeVars {
	vars := WelcomeVars{}
	vars.Rows = FetchStories(db, order, domain)
	return &vars
}
