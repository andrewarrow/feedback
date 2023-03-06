package router

import (
	"fmt"
	"html"
	"html/template"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Comment struct {
	Guid       string
	Ago        string
	Timestamp  string
	Username   string
	Body       template.HTML
	RawBody    string
	StoryGuid  string
	StoryTitle string
}

func FetchComments(db *sqlx.DB, storyId int64) []*Comment {
	items := []*Comment{}
	rows, err := db.Queryx("SELECT * FROM comments where story_id = $1 ORDER BY created_at desc limit 30", storyId)
	if err != nil {
		return items
	}
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		comment := commentFromMap(m)
		items = append(items, comment)
	}
	return items
}

func FetchComment(db *sqlx.DB, guid string) *Comment {
	rows, err := db.Queryx("SELECT * FROM comments where guid = $1", guid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	rows.Next()
	m := make(map[string]any)
	rows.MapScan(m)
	if len(m) == 0 {
		return nil
	}
	comment := commentFromMap(m)
	return comment
}

func commentFromMap(m map[string]any) *Comment {
	c := Comment{}
	c.Guid = fmt.Sprintf("%s", m["guid"])
	c.StoryGuid = fmt.Sprintf("%s", m["story_guid"])
	c.RawBody = fmt.Sprintf("%s", m["body"])
	c.Username = fmt.Sprintf("%s", m["username"])
	body := strings.Replace(html.EscapeString(c.RawBody), "\n", "<br/>", -1)
	c.Body = template.HTML(body + "<br/>")

	c.Timestamp, c.Ago = FixTime(m)
	return &c
}
