package controllers

import "github.com/gin-gonic/gin"
import "net/http"

func UsersIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"name":  "name",
	})

}
func UsersShow(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"flash": "",
		"name":  "name",
	})

}
