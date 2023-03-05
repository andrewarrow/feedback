package router

import (
	"fmt"
	"net/http"

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
	if path == "/users/" {
		return false
	}
	return true
}

func (r *Router) LookupUser(guid string) *models.User {
	if guid == "" {
		return nil
	}
	rows, _ := r.Db.Queryx("SELECT * FROM users where guid=$1", guid)
	m := make(map[string]any)
	rows.Next()
	rows.MapScan(m)
	if len(m) == 0 {
		return nil
	}
	user := models.User{}
	user.Username = fmt.Sprintf("%s", m["username"])
	return &user
}

func (r *Router) LookupUsername(username string) *models.User {
	if username == "" {
		return nil
	}
	rows, _ := r.Db.Queryx("SELECT * FROM users where username=$1", username)
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
