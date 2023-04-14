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
		headers := "-H \"Authorization: Bearer token\" -H \"Content-Type: json\""
		payload := m.CurlCreatePayload()
		fmt.Printf("curl -XPOST %s http://localhost:8080/%s -d %s\n", headers, route.Root, payload)
		fmt.Println("```")
		fmt.Println("")
		fmt.Println("### Example response")
		fmt.Println("```json")
		response := `{"thing": [{"more": "this"}]}`
		fmt.Println(response)
		fmt.Println("```")

		fmt.Println("")
	}
}
