package controllers

import (
	"net/http"

	"github.com/andrewarrow/feedback/util"
	"github.com/gin-gonic/gin"
)

func LegalGdpr(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("gdpr_ok", "cookies, yes", 0, "/", host, false, false)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

func LegalPrivacy(c *gin.Context) {
	c.HTML(http.StatusOK, "privacy.tmpl", gin.H{
		"flash": "",
	})

}
func LegalTerms(c *gin.Context) {
	c.HTML(http.StatusOK, "terms.tmpl", gin.H{
		"flash": "",
	})

}