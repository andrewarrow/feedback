package controllers

import (
	"github.com/andrewarrow/feedback/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ApiVersion(c *gin.Context) {
	ap := api.ApiResponse{}
	ap.Version = "1.0.0"
	ap.Items = []interface{}{"test1", "test2"}
	ap.SentAt = time.Now().Unix()
	c.JSON(http.StatusOK, ap)

}
