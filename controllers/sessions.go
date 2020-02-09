package controllers

import "github.com/gin-gonic/gin"
import "strings"
import "fmt"
import "github.com/andrewarrow/feedback/models"
import "github.com/andrewarrow/feedback/util"
import "github.com/tjarratt/babble"
import "net/http"

var babbler = babble.NewBabbler()

func SessionsNew(c *gin.Context) {
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
	host := util.AllConfig.Http.Host

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") || len(email) < 7 {
		flash = "not valid email"
	} else {
		user.Email = email
		user.Flavor = "user"
		sql := fmt.Sprintf("SELECT email, flavor from users where email=:email and phrase=SHA1(:phrase)")
		rows, err := Db.NamedQuery(sql, map[string]interface{}{"email": email, "phrase": password})
		if err != nil {
			flash = err.Error()
		} else {
			if rows.Next() {
				rows.StructScan(&user)
				c.SetCookie("user", user.Encode(), 3600, "/", host, false, false)
			} else {
				babbler.Count = 4
				phrase := babbler.Babble()
				m := map[string]interface{}{"email": email, "phrase": phrase, "flavor": "user"}
				_, err = Db.NamedExec(`INSERT INTO users (email, phrase, flavor) 
values (:email, SHA1(:phrase), :flavor)`, m)
				if err != nil {
					flash = "was not able to login"
				} else {
					c.SetCookie("user", user.Encode(), 3600, "/", host, false, false)
				}
			}
		}
	}
	c.SetCookie("flash", flash, 3600, "/", host, false, false)
	c.Redirect(http.StatusMovedPermanently, "/")
	c.Abort()
}
func SessionsDestroy(c *gin.Context) {
	host := util.AllConfig.Http.Host
	c.SetCookie("user", "", 3600, "/", host, false, false)

	c.Redirect(http.StatusMovedPermanently, "/")
	c.Abort()
}
