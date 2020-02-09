package controllers

import "github.com/gin-gonic/gin"
import "net/http"
import "github.com/jmoiron/sqlx"
import "github.com/andrewarrow/feedback/util"
import "github.com/andrewarrow/feedback/models"

var Db *sqlx.DB
var flash = ""

func BeforeAll(c *gin.Context) {
	flash, _ = c.Cookie("flash")
	host := util.AllConfig.Http.Host
	c.SetCookie("flash", "", 3600, "/", host, false, false)
}

func SetFlash(s string, c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("flash", s, 3600, "/", host, false, false)
}

func WelcomeIndex(c *gin.Context) {
	BeforeAll(c)
	json, _ := c.Cookie("user")
	user := models.DecodeUser(json)

	if user == nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"user":  nil,
			"flash": flash,
		})
		return
	}

	c.HTML(http.StatusOK, "homepage.tmpl", gin.H{
		"user":  user,
		"flash": flash,
	})
}
