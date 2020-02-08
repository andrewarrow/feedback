package controllers

import "github.com/gin-gonic/gin"
import "github.com/andrewarrow/feedback/models"
import "net/http"

func AdminUsersIndex(c *gin.Context) {
	users, err := models.SelectUsers(Db)

	c.HTML(http.StatusOK, "users__index.tmpl", gin.H{
		"flash": err,
		"users": users,
		"name":  "name",
	})

}
func AdminUsersShow(c *gin.Context) {
	c.HTML(http.StatusOK, "users__show.tmpl", gin.H{
		"flash": "",
		"user":  map[string]string{},
		"name":  "name",
	})

}
