package app

import (
  "github.com/andrewarrow/feedback/router"
)

func Z{{homeducky}}(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleZ{{homeducky}}Index(c)
		return
	}
	c.NotFound = true
}

func handleZ{{homeducky}}Index(c *router.Context) {

	send := map[string]any{}
  c.SendContentInLayout("welcome.html", send, 200)
}
    
