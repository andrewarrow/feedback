package router

import (
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/util"
)

func handleComments(c *Context, second, third string) {
	if second == "" {
		c.notFound = true
	} else if third != "" {
		c.notFound = true
	} else {
		if c.method == "POST" {
			postComment(c, second)
		} else {
			showComment(c, second)
		}
	}
}

func showComment(c *Context, second string) {
	comment := FetchComment(c.db, second)
	if comment == nil {
		c.notFound = true
		return
	}
	c.title = comment.RawBody
	if len(c.title) > 80 {
		c.title = c.title[0:80] + "..."
	}
	story := FetchStory(c.db, comment.StoryGuid)
	if story == nil {
		c.notFound = true
		return
	}
	comment.StoryTitle = story.Title
	if len(comment.StoryTitle) > 40 {
		comment.StoryTitle = comment.StoryTitle[0:40] + "..."
	}
	c.SendContentInLayout("comments_show.html", comment, 200)
	return
}

func postComment(c *Context, second string) {
	body := strings.TrimSpace(c.request.FormValue("body"))
	returnPath := "/stories/" + second + "/"
	if len(body) < 10 {
		setFlash(c, "body too short.")
		http.Redirect(c.writer, c.request, returnPath, 302)
		return
	}

	guid := util.PseudoUuid()
	story := FetchStory(c.db, second)
	if story == nil {
		c.notFound = true
		return
	}

	tx := c.db.MustBegin()
	tx.Exec("insert into comments (body, guid, username, story_id, story_guid) values ($1, $2, $3, $4, $5)", body, guid, c.user.Username, story.Id, story.Guid)
	tx.Exec("update stories set comments=comments+1 where id=$1", story.Id)
	tx.Commit()
	http.Redirect(c.writer, c.request, returnPath, 302)
}
