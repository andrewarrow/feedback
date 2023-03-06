package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/andrewarrow/feedback/models"
)

func (r *Router) IsUserRequired(path string, method string) bool {
	//fmt.Println(path, method)
	if path == "/sessions/new/" {
		return false
	}
	if path == "/sessions/" {
		return false
	}
	if path == "/fresh/" {
		return false
	}
	if path == "/about/" {
		return false
	}
	if strings.HasPrefix(path, "/stories/") && method == "GET" {
		return false
	}
	if strings.HasPrefix(path, "/comments/") && method == "GET" {
		return false
	}
	if strings.HasPrefix(path, "/users/") {
		return false
	}
	return true
}

func (c *Context) IsAdmin() bool {
	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "*" {
		return true
	}
	return c.user.Guid == adminUser
}

func (r *Router) LookupUser(guid string) *models.User {
	if guid == "" {
		return nil
	}
	rows, err := r.Db.Queryx("SELECT * FROM users where guid=$1", guid)
	if err != nil {
		return nil
	}
	defer rows.Close()
	m := make(map[string]any)
	rows.Next()
	rows.MapScan(m)
	if len(m) == 0 {
		return nil
	}
	user := models.User{}
	user.Username = fmt.Sprintf("%s", m["username"])
	user.Guid = guid
	return &user
}

func (r *Router) LookupUsername(username string) *models.User {
	if username == "" {
		return nil
	}
	rows, err := r.Db.Queryx("SELECT * FROM users where username=$1", username)
	if err != nil {
		return nil
	}
	defer rows.Close()
	m := make(map[string]any)
	rows.Next()
	rows.MapScan(m)
	if len(m) == 0 {
		return nil
	}
	user := models.User{}
	user.Username = fmt.Sprintf("%s", m["username"])
	user.Timestamp, user.Ago = FixTime(m)
	return &user
}

func DestroyFlash(writer http.ResponseWriter) {
	cookie := http.Cookie{}
	cookie.MaxAge = 0
	cookie.Name = "flash"
	cookie.Value = ""
	cookie.Path = "/"
	http.SetCookie(writer, &cookie)
}
