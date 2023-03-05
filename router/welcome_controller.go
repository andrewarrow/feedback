package router

import (
	"fmt"
	"html"
	"html/template"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/models"
	"github.com/jmoiron/sqlx"
	"github.com/xeonx/timeago"
)

type Story struct {
	Url       string
	Title     string
	Guid      string
	Ago       string
	Timestamp string
	Username  string
	HasUrl    bool
	Domain    string
	Body      template.HTML
	Id        int64
	Comments  int64
	AddS      string
}

type WelcomeVars struct {
	Rows []*Story
}

const layout = "2006-01-02 15:04:05"

func WelcomeIndexVars(db *sqlx.DB, location *time.Location) *WelcomeVars {
	vars := WelcomeVars{}
	vars.Rows = []*Story{}
	rows, err := db.Queryx("SELECT * FROM stories ORDER BY created_at desc limit 30")
	if err != nil {
		return &vars
	}
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		story := storyFromMap(m)
		vars.Rows = append(vars.Rows, story)
	}
	return &vars
}

func storyFromMap(m map[string]any) *Story {
	story := Story{}
	story.Title = models.RemoveMostNonAlphanumeric(fmt.Sprintf("%s", m["title"]))
	story.Url = fmt.Sprintf("%s", m["url"])
	story.Guid = fmt.Sprintf("%s", m["guid"])
	story.Username = fmt.Sprintf("%s", m["username"])
	story.Id = m["id"].(int64)
	story.Comments = m["comments"].(int64)
	if story.Comments != 1 {
		story.AddS = "s"
	}
	body := fmt.Sprintf("%s", m["body"])
	body = strings.Replace(html.EscapeString(body), "\n", "<br/>", -1)
	story.Body = template.HTML(body + "<br/><br/>")
	if story.Url != "" {
		story.HasUrl = true
		tokens := strings.Split(story.Url, "/")
		if len(tokens) > 2 {
			tokens = strings.Split(tokens[2], ".")
			if len(tokens) == 3 {
				tokens = tokens[1:]
			}
			story.Domain = strings.Join(tokens, ".")
		}
	}

	tm := m["created_at"].(time.Time)
	tm = tm.Add(time.Hour * 8)
	story.Timestamp = fmt.Sprintf("%s", tm)
	story.Ago = timeago.English.Format(tm)
	return &story
}
