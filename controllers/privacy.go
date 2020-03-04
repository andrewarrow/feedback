package controllers

import (
	"net/http"

	"github.com/andrewarrow/feedback/util"
	"github.com/gin-gonic/gin"
)

func GdprAsk(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("gdpr_ok", "cookies, yes", 0, "/", host, false, false)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
