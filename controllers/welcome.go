package controllers

import "github.com/gin-gonic/gin"
import "net/http"
import "github.com/jmoiron/sqlx"

//import "github.com/andrewarrow/feedback/models"

var Db *sqlx.DB

func WelcomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"user":  nil,
		"flash": "",
		"name":  "name",
	})

}
