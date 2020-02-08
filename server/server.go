package server

import "github.com/gin-gonic/gin"
import "net/http"
import "fmt"
import "time"
import "github.com/andrewarrow/feedback/util"
import "github.com/andrewarrow/feedback/persist"
import "github.com/jmoiron/sqlx"

var db *sqlx.DB

func Serve() {
	prefix := util.AllConfig.Path.Prefix

	db = persist.Connection()
	router := gin.Default()

	router.Static("/assets", prefix+"assets")
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "index")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	AddTemplates(router, prefix)
	go router.Run(fmt.Sprintf(":%d", util.AllConfig.Http.Port))

	for {
		time.Sleep(time.Second)
	}

}
