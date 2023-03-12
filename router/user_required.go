package router

import (
	"errors"
	"fmt"
	"sort"
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

func UserRequiredPathsSorted(m map[string]*UserRequired) []string {
	items := []string{}
	for k, _ := range m {
		items = append(items, k)
	}
	sort.Strings(items)
	return items
}

func (ur *UserRequired) ShouldRequire(urPath, path string, method string) (bool, error) {
	pathMatch := false
	if urPath != path && !strings.HasPrefix(path, urPath) {
		return false, errors.New("wrong area")
	}
	fmt.Println(urPath, path, method)

	if ur.MatchBy == "==" {
		pathMatch = urPath == path
	} else if ur.MatchBy == "prefix" {
		pathMatch = strings.HasPrefix(path, urPath)
	}

	if pathMatch && ur.Method == "*" {
		return false, nil
	}

	if pathMatch && ur.Method == method {
		return false, nil
	}

	return true, nil
}
