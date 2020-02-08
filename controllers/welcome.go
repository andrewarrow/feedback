package controllers

import "github.com/gin-gonic/gin"
import "net/http"
import "github.com/jmoiron/sqlx"
import "fmt"

import "github.com/andrewarrow/feedback/models"

var Db *sqlx.DB

func WelcomeIndex(c *gin.Context) {
	json, _ := c.Cookie("user")
	fmt.Println("11111111111111", json)
	user := models.DecodeUser(json)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"user":  user,
		"flash": "",
		"name":  "name",
	})

}
