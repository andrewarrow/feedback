package gogen

import (
	"fmt"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func MakeRoutes(routes []*models.Route) {
	for _, route := range routes {
		fmt.Println(route.Root)
		output := route.Generate(route.Root)
		fmt.Println(output)
	}
}

func MakeMarkDown(s *router.FeedbackSite) {

	fmt.Println("")
	fmt.Println("")
	for _, route := range s.Routes {
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
		fmt.Println("### Example curls")
		fmt.Println("```")
		modelString := util.Unplural(route.Root)
		m := s.FindModel(modelString)
		fmt.Println(m)
		fmt.Println("curl http://localhost:8080/thing")
		fmt.Println("```")

		fmt.Println("")
	}
}
