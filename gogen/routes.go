package gogen

import (
	"fmt"

	"github.com/andrewarrow/feedback/models"
)

func MakeRoutes(routes []*models.Route) {
	for _, route := range routes {
		fmt.Println(route.Root)
		output := route.Generate(route.Root)
		fmt.Println(output)
	}

	fmt.Println("")
	fmt.Println("")
	for _, route := range routes {
		fmt.Println("# " + route.Root)
		fmt.Println("")
		fmt.Println("```")
		for _, path := range route.Paths {
			more := ""
			if path.Second == "*" {
				more = "/:guid"
			} else if path.Third == "*" {
				more = "/" + path.Second + "/:something"
			}
			fmt.Printf("% 6s /%s%s\n", path.Verb, route.Root, more)
		}
		fmt.Println("```")
		fmt.Println("")
	}
}
