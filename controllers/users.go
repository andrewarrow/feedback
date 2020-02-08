package controllers

import "github.com/gin-gonic/gin"
import "net/http"

func UsersIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "users__index.tmpl", gin.H{
		"flash": "",
		"users": []string{},
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
