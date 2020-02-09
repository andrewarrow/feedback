package controllers

import "github.com/gin-gonic/gin"
import "github.com/andrewarrow/feedback/models"
import "net/http"

var user *models.User

func ValidAdminUser(c *gin.Context) bool {
	json, _ := c.Cookie("user")
	user = models.DecodeUser(json)
	if user == nil || user.Flavor != "admin" {
		SetFlash("you need to login", c)
		c.Redirect(http.StatusFound, "/sessions/new")
		c.Abort()
		return false
	}
	return true
}
func UsersIndex(c *gin.Context) {
	if !ValidAdminUser(c) {
		return
	}
	users, err := models.SelectUsers(Db)

	c.HTML(http.StatusOK, "users__index.tmpl", gin.H{
		"flash": err,
		"users": users,
		"name":  "name",
	})

}
func UsersShow(c *gin.Context) {
	c.HTML(http.StatusOK, "users__show.tmpl", gin.H{
		"flash": "",
		"user":  map[string]string{},
		"name":  "name",
	})

}
