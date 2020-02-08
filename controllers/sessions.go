package controllers

import "github.com/gin-gonic/gin"

//import "github.com/andrewarrow/feedback/models"
import "net/http"

func SessionsNew(c *gin.Context) {
	c.HTML(http.StatusOK, "sessions__new.tmpl", gin.H{
		"flash": "",
		"name":  "name",
	})

}
