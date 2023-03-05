package router

import (
	"fmt"
	"html"
	"html/template"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/jmoiron/sqlx"
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
	Points    int64
}

type WelcomeVars struct {
	Rows []*Story
}

const layout = "2006-01-02 15:04:05"

func WelcomeIndexVars(db *sqlx.DB, order string) *WelcomeVars {
	vars := WelcomeVars{}
	vars.Rows = []*Story{}
	rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM stories ORDER BY %s limit 30", order))
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
	story.Points = m["points"].(int64)
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

	story.Timestamp, story.Ago = FixTime(m)
	return &story
}
