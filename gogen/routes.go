package gogen

import (
	"fmt"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func MakeRoutes(routes []*models.Route, modelString string) {
	for _, route := range routes {
		if route.Root != modelString {
			continue
		}
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
		if path.Second != "" && path.Third == "" {
			more = "/:guid"
		} else if path.Second == "*" && path.Third != "" {
			more = fmt.Sprintf("/%s/:guid", path.Third)
		} else if path.Third == "*" {
			more = "/" + path.Second + "/:something"
		}
		if path.Params != "" {
			more = more + "?" + path.Params
		}
		fmt.Printf("% 6s /%s%s\n", path.Verb, route.Root, more)
	}
}

func deleteRoute(root, headers string, m *models.Model) {
	fmt.Println("```")
	fmt.Printf("curl -XDELETE %s http://localhost:8080/%s\n", headers, root)
	fmt.Println("```")
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

func show(root, headers string, m *models.Model, guid string) {
	fmt.Println("```")
	fmt.Printf("curl -XGET %s http://localhost:8080/%s/%s\n", headers, root, guid)
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
		m := s.FindModelOrDynamic(util.FixForDash(modelString))
		headers := "-H \"Authorization: Bearer token\" -H \"Content-Type: json\""
		fmt.Println("```")
		fmt.Println("")
		for _, path := range route.Paths {
			fmt.Printf("## %s curl\n", path.Verb)
			if path.Verb == "GET" && path.Second == "" {
				index(route.Root, headers, m)
			} else if path.Verb == "GET" && path.Second != "" {
				show(route.Root, headers, m, path.Second)
			} else if path.Verb == "PUT" {
				put(route.Root, headers, m)
			} else if path.Verb == "POST" {
				post(route.Root, headers, m)
			} else if path.Verb == "DELETE" {
				deleteRoute(route.Root, headers, m)
			}
			fmt.Println("")
			if path.Response == "bool" {
				if path.Verb == "POST" {
					fmt.Println("### Payload")
					fmt.Println("```json")
					fmt.Println(m.CurlSingleResponseNoWrapper())
					fmt.Println("```")
					fmt.Println("")
				}
				fmt.Println("### Response")
				fmt.Println("```json")
				fmt.Println(`{"info": "ok"}`)
				fmt.Println("```")
				fmt.Println("")
				fmt.Println("```json")
				fmt.Println(`{"info": "stripe api down"}`)
				fmt.Println("```")
				fmt.Println("")
			} else if path.Response != "" {
				m := s.FindDynamic(util.FixForDash(path.Response))
				fmt.Println("### Single response")
				fmt.Println("```json")
				fmt.Println(m.CurlSingleResponse())
				fmt.Println("```")
				fmt.Println("")
			}
		}
		fmt.Println("### Single response")
		fmt.Println("```json")
		fmt.Println(m.CurlSingleResponse())
		fmt.Println("```")
		fmt.Println("")
		fmt.Println("### List response")
		fmt.Println("```json")
		fmt.Println(m.CurlListResponse())
		fmt.Println("```")
		fmt.Println("")
	}
}
