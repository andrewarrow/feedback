package router

import (
	"fmt"
	"html"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/util"
	"github.com/jmoiron/sqlx"
	"github.com/xeonx/timeago"
)

func handleComments(c *Context, second, third string) {
	if second == "" {
		c.notFound = true
	} else if third != "" {
		c.notFound = true
	} else {
		body := strings.TrimSpace(c.request.FormValue("body"))
		returnPath := "/stories/" + second + "/"
		if len(body) < 10 {
			setFlash(c, "body too short.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}

		guid := util.PseudoUuid()
		story := FetchStory(c.db, second)

		tx := c.db.MustBegin()
		tx.Exec("insert into comments (body, guid, username, story_id) values ($1, $2, $3, $4)", body, guid, c.user.Username, story.Id)
		tx.Exec("update stories set comments=comments+1 where id=$1", story.Id)
		tx.Commit()
		http.Redirect(c.writer, c.request, returnPath, 302)
	}
}

type Comment struct {
	Guid      string
	Ago       string
	Timestamp string
	Username  string
	Body      template.HTML
}

func FetchComments(db *sqlx.DB, storyId int64) []*Comment {
	items := []*Comment{}
	rows, err := db.Queryx("SELECT * FROM comments where story_id = $1 ORDER BY created_at desc limit 30", storyId)
	if err != nil {
		return items
	}
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		comment := commentFromMap(m)
		items = append(items, comment)
	}
	return items
}

func commentFromMap(m map[string]any) *Comment {
	c := Comment{}
	c.Guid = fmt.Sprintf("%s", m["guid"])
	body := fmt.Sprintf("%s", m["body"])
	body = strings.Replace(html.EscapeString(body), "\n", "<br/>", -1)
	c.Body = template.HTML(body + "<br/><br/>")

	tm := m["created_at"].(time.Time)
	tm = tm.Add(time.Hour * 8)
	c.Timestamp = fmt.Sprintf("%s", tm)
	c.Ago = timeago.English.Format(tm)
	return &c
}
