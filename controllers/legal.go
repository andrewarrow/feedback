package controllers

import (
	"net/http"

	"github.com/andrewarrow/feedback/util"
	"github.com/gin-gonic/gin"
)

func LegalGdpr(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("gdpr_ok", "cookies", 2147483647, "/", "localhost", false, true)

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}

func LegalPrivacy(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.HTML(http.StatusOK, "privacy.tmpl", gin.H{
		"flash": "",
		"name":  "feedback", // hint: change me
		"host":  host,
	})

}
func LegalTerms(c *gin.Context) {
	host := util.AllConfig.Http.Host

	c.HTML(http.StatusOK, "terms.tmpl", gin.H{
		"flash": "",
		"name":  "feedback", // hint: change me
		"host":  host,
	})

}
