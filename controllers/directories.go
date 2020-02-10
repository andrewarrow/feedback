package controllers

import "github.com/gin-gonic/gin"
import "github.com/andrewarrow/feedback/models"
import "net/http"

func DirectoriesIndex(c *gin.Context) {
	if !BeforeAll("user", c) {
		return
	}
	users, err := models.SelectUsers(Db)

	c.HTML(http.StatusOK, "users__index.tmpl", gin.H{
		"flash": err,
		"users": users,
		"name":  "name",
	})

}
func DirectoriesDownload(c *gin.Context) {
}
func DirectoriesNameIndex(c *gin.Context) {
}
func DirectoriesDownloadExtra(c *gin.Context) {
}
