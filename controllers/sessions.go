package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/gin-gonic/gin"
	"github.com/tjarratt/babble"
)

var babbler = babble.NewBabbler()

func SessionsNew(c *gin.Context) {
	if !BeforeAll("", c) {
		return
	}
	c.HTML(http.StatusOK, "sessions__new.tmpl", gin.H{
		"flash": "",
		"name":  "name",
	})

}
func SessionsCreate(c *gin.Context) {
	user := models.User{}
	email := c.PostForm("email")
	password := c.PostForm("password")
	flash := ""

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") || len(email) < 7 {
		flash = "not valid email"
	} else {
		user.Email = email
		user.Flavor = "user"
		sql := fmt.Sprintf("SELECT id, email, flavor from users where email=:email and phrase=SHA1(:phrase)")
		rows, err := Db.NamedQuery(sql, map[string]interface{}{"email": email, "phrase": password})
		if err != nil {
			flash = err.Error()
		} else {
			if rows.Next() {
				rows.StructScan(&user)
				c.SetCookie("user", user.Encode(), 3600*24*365, "/", "localhost", false, true)
			} else {
				babbler.Count = 4
				phrase := babbler.Babble()
				fmt.Println(phrase)
				m := map[string]interface{}{"email": email, "phrase": phrase, "flavor": "user"}
				_, err = Db.NamedExec(`INSERT INTO users (email, phrase, flavor) 
values (:email, SHA1(:phrase), :flavor)`, m)
				if err != nil {
					flash = "was not able to login"
				} else {
					c.SetCookie("user", user.Encode(), 3600, "/", "localhost", false, true)
				}
			}
		}
	}
	c.SetCookie("flash", flash, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
func SessionsDestroy(c *gin.Context) {
	c.SetCookie("user", "", 3600, "/", "localhost", false, true)

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
