package router

import (
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/util"
	"github.com/jmoiron/sqlx"
)

type StoryShow struct {
	Story    *Story
	Comments []*Comment
}

func handleStories(c *Context, second, third string) {
	if second == "" {
		handleStoriesIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		if second == "new" {
			c.SendContentInLayout("stories_new.html", nil, 200)
			return
		} else if second != "" {
			storyShow := StoryShow{}
			storyShow.Story = FetchStory(c.db, second)
			if storyShow.Story == nil {
				c.notFound = true
				return
			}
			storyShow.Comments = FetchComments(c.db, storyShow.Story.Id)
			c.SendContentInLayout("stories_show.html", storyShow, 200)
			return
		}
		c.notFound = true
	}
}

func FetchStory(db *sqlx.DB, guid string) *Story {
	rows, err := db.Queryx("select * from stories where guid=$1", guid)
	if err != nil {
		return nil
	}
	rows.Next()
	m := make(map[string]any)
	rows.MapScan(m)
	return storyFromMap(m)
}

func handleStoriesIndex(c *Context) {
	if c.method == "POST" {
		title := strings.TrimSpace(c.request.FormValue("title"))
		url := strings.TrimSpace(c.request.FormValue("url"))
		body := strings.TrimSpace(c.request.FormValue("body"))
		returnPath := "/stories/new/"
		if len(title) < 10 {
			setFlash(c, "title too short.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}
		if len(title) > 140 {
			setFlash(c, "title too long.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}
		if body == "" && url == "" {
			setFlash(c, "body or url required.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}
		if url != "" {
			if len(url) < 13 ||
				!(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
				setFlash(c, "url too short.")
				http.Redirect(c.writer, c.request, returnPath, 302)
				return
			}
		}
		if body != "" && len(body) < 10 {
			setFlash(c, "body too short.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}
		guid := util.PseudoUuid()
		if url != "" {
			c.db.Exec("insert into stories (title, url, guid, username) values ($1, $2, $3, $4)", title, url, guid, c.user.Username)
		} else {
			c.db.Exec("insert into stories (title, body, guid, username) values ($1, $2, $3, $4)", title, body, guid, c.user.Username)
		}
		http.Redirect(c.writer, c.request, "/", 302)
		return
	}
	c.notFound = true
}
