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

func printRoutes(route *models.Route) {
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
}

func post(root, headers string, m *models.Model) {
	fmt.Println("```")
	payload := m.CurlPostPayload()
	fmt.Printf("curl -XPOST %s http://localhost:8080/%s -d %s\n", headers, root, payload)
	fmt.Println("```")
}
func index(root, headers string, m *models.Model) {
	fmt.Println("```")
	fmt.Printf("curl -XGET %s http://localhost:8080/%s\n", headers, root)
	fmt.Println("```")
}

func put(root, headers string, m *models.Model) {
	fmt.Println("```")
	payload := m.CurlPutPayload()
	fmt.Printf("curl -XPUT %s http://localhost:8080/%s/%s -d %s\n", headers, root, util.PseudoUuid(), payload)
	fmt.Println("```")
}

func show(root, headers string, m *models.Model) {
	fmt.Println("```")
	fmt.Printf("curl -XGET %s http://localhost:8080/%s/%s\n", headers, root, util.PseudoUuid())
	fmt.Println("```")
}

func MakeMarkDown(s *router.FeedbackSite, modelString string) {

	fmt.Println("")
	fmt.Println("")
	for _, route := range s.Routes {
		if util.Unplural(route.Root) != modelString {
			continue
		}

		printRoutes(route)
		m := s.FindModel(modelString)
		headers := "-H \"Authorization: Bearer token\" -H \"Content-Type: json\""
		fmt.Println("```")
		fmt.Println("")
		fmt.Println("### Example curls")
		fmt.Println("")
		post(route.Root, headers, m)
		fmt.Println("")
		put(route.Root, headers, m)
		fmt.Println("")
		index(route.Root, headers, m)
		fmt.Println("")
		show(route.Root, headers, m)
		fmt.Println("")
		fmt.Println("### Example response")
		fmt.Println("```json")
		fmt.Println(m.CurlResponse())
		fmt.Println("```")

		fmt.Println("")
	}
}
