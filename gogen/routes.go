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

func MakeMarkDown(s *router.FeedbackSite, modelString string) {

	fmt.Println("")
	fmt.Println("")
	for _, route := range s.Routes {
		if util.Unplural(route.Root) != modelString {
			continue
		}

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
		m := s.FindModel(modelString)
		headers := "-H \"Authorization: Bearer token\" \\ \n  -H \"Content-Type: json\" \\ \n"
		fmt.Println("```")
		fmt.Println("")
		fmt.Println("### Example curls")
		fmt.Println("```")
		payload := m.CurlPostPayload()
		fmt.Printf("curl -XPOST %s http://localhost:8080/%s -d %s\n", headers, route.Root, payload)
		fmt.Println("```")
		fmt.Println("")
		fmt.Println("```")
		payload = m.CurlPutPayload()
		fmt.Printf("curl -XPUT %s http://localhost:8080/%s -d %s\n", headers, route.Root, payload)
		fmt.Println("```")
		fmt.Println("")
		fmt.Println("### Example response")
		fmt.Println("```json")
		fmt.Println(m.CurlResponse())
		fmt.Println("```")

		fmt.Println("")
	}
}
