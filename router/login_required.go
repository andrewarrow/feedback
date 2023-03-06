package router

import (
	"strings"
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
