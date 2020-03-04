package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andrewarrow/feedback/controllers"
	"github.com/andrewarrow/feedback/persist"
	"github.com/andrewarrow/feedback/util"
	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	prefix := util.AllConfig.Path.Prefix

	controllers.Db = persist.Connection()
	router := gin.Default()

	router.Static("/assets", prefix+"assets")
	router.GET("/", controllers.WelcomeIndex)
	router.GET("/gdpr", controllers.GdprAsk)
	inboxes := router.Group("/inboxes")
	inboxes.GET("/", controllers.InboxesIndex)
	users := router.Group("/users")
	users.GET("/", controllers.UsersIndex)
	user := router.Group("/user")
	user.GET("/:id", controllers.UsersShow)
	sessions := router.Group("/sessions")
	sessions.GET("/new", controllers.SessionsNew)
	sessions.POST("/", controllers.SessionsCreate)
	sessions.POST("/destroy", controllers.SessionsDestroy)

	admin := router.Group("/admin")
	users = admin.Group("/users")
	users.GET("/", controllers.AdminUsersIndex)
	user = admin.Group("/user")
	user.GET("/:id", controllers.AdminUsersShow)

	guid := util.AllConfig.Directories.Guid
	if guid != "" {
		router.GET("/"+guid+"/", controllers.DirectoriesIndex)
		router.GET("/"+guid+"/:name", controllers.DirectoriesDownload)
		router.GET("/"+guid+"/:name/", controllers.DirectoriesNameIndex)
		router.GET("/"+guid+"/:name/:extra", controllers.DirectoriesDownloadExtra)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	AddTemplates(router, prefix)
	go router.Run(fmt.Sprintf(":%s", port))

	for {
		time.Sleep(time.Second)
	}

}
