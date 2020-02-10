package controllers

import "github.com/gin-gonic/gin"
import "github.com/andrewarrow/feedback/models"
import "net/http"
import "strconv"

func InboxesIndex(c *gin.Context) {
	if !BeforeAll("admin", c) {
		return
	}
	items, err := models.SelectInboxes(Db)

	selected := ""
	if len(items) > 0 {
		index := c.DefaultQuery("i", "0")
		i, _ := strconv.Atoi(index)
		selected = items[i].Body
	}

	c.HTML(http.StatusOK, "inboxes__index.tmpl", gin.H{
		"flash":    err,
		"selected": selected,
		"user":     user,
		"items":    items,
	})

}
