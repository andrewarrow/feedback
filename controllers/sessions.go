package controllers

import "github.com/gin-gonic/gin"

import "github.com/andrewarrow/feedback/models"
import "github.com/andrewarrow/feedback/util"
import "net/http"

func SessionsNew(c *gin.Context) {
	c.HTML(http.StatusOK, "sessions__new.tmpl", gin.H{
		"flash": "",
		"name":  "name",
	})

}
func SessionsCreate(c *gin.Context) {
	user := models.User{}
	user.Email = "wfwe"

	host := util.AllConfig.Http.Host
	c.SetCookie("user", user.Encode(), 3600, "/", host, false, false)

	c.Redirect(http.StatusMovedPermanently, "/")
	c.Abort()
}
