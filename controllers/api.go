package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ApiReponse struct {
	Version string        `json:"version"`
	Items   []interface{} `json:"items"`
	SentAt  int64         `json:"sent_at"`
}

func ApiVersion(c *gin.Context) {
	ap := ApiReponse{}
	ap.Version = "1.0.0"
	ap.Items = []interface{}{"test1", "test2"}
	ap.SentAt = time.Now().Unix()
	c.JSON(http.StatusOK, ap)

}
