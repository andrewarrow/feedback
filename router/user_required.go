package router

import (
	"strings"
)

type UserRequired struct {
	Method  string
	MatchBy string
}

func NewUserRequired(method, matchBy string) *UserRequired {
	ur := UserRequired{}
	ur.Method = method
	ur.MatchBy = matchBy
	return &ur
}

func (ur *UserRequired) ShouldNotRequire(urPath, path string, method string) bool {
	pathMatch := false
	if urPath != path && !strings.HasPrefix(path, urPath) {
		return false
	}

	if ur.MatchBy == "==" {
		pathMatch = urPath == path
	} else if ur.MatchBy == "prefix" {
		pathMatch = strings.HasPrefix(path, urPath)
	}

	if pathMatch && ur.Method == "*" {
		return true
	}

	if pathMatch && ur.Method == method {
		return true
	}

	return false
}
