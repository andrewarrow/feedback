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
}
