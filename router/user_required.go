package router

import (
	"strings"
)

type UserRequired struct {
	Path    string
	Method  string
	MatchBy string
}

func NewUserRequired(path, method, matchBy string) *UserRequired {
	ur := UserRequired{}
	ur.Path = path
	ur.Method = method
	ur.MatchBy = matchBy
	return &ur
}

func (ur *UserRequired) ShouldRequire(path string, method string) bool {
	pathMatch := false
	//fmt.Println(urPath, path, method)

	if ur.MatchBy == "==" {
		pathMatch = ur.Path == path
	} else if ur.MatchBy == "prefix" {
		pathMatch = strings.HasPrefix(path, ur.Path)
	}

	if pathMatch && ur.Method == "*" {
		return true
	}

	if pathMatch && ur.Method == method {
		return true
	}

	return false
}
