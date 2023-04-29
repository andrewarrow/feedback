package router

import (
	"fmt"
	"strings"
)

func (c *Context) FunctionToRun(route string) func(*Context, string, string) {
	newContext := Context{}

	tokens := strings.Split(route, "?")
	noParams := tokens[0]
	fmt.Println("noParams", noParams)
	newContext.tokens = strings.Split(noParams, "/")
	first := newContext.tokens[1]

	return c.router.pathFuncToRun(first)
}
