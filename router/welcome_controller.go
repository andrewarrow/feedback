package router

import (
	"encoding/json"
	"fmt"
)

func handleWelcome(c *Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleWelcomeIndex(c)
		return
	}
	c.NotFound = true
}

func handleWelcomeIndex(c *Context) {
	list := getData()

	colAttributes := map[int]string{}
	colAttributes[0] = "w-1/2"

	m := map[string]any{}
	headers := []string{"name", "age", "species", "home_planet", "language", "occupation"}

	params := map[string]any{}
	m["headers"] = headers
	m["cells"] = c.MakeCells(list, headers, params, "_welcome")
	m["col_attributes"] = colAttributes

	send := map[string]any{}
	send["bottom"] = c.Template("table_show.html", m)
	c.SendContentInLayout("generic_top_bottom.html", send, 200)
}

func getData() []any {
	data := `
{
  "aliens": [
    {
      "name": "Zog",
      "age": 150,
      "species": "Xenon",
      "home_planet": "Zeta-7",
      "language": "Zorgon",
      "occupation": "Astrobiologist"
    },
    {
      "name": "Luna",
      "age": 200,
      "species": "Lunarian",
      "home_planet": "Moon",
      "language": "Lunar",
      "occupation": "Quantum Physicist"
    },
    {
      "name": "Glimmer",
      "age": 75,
      "species": "Nebulite",
      "home_planet": "Nebula-9",
      "language": "Stellar",
      "occupation": "Astroengineer"
    },
    {
      "name": "Xylon",
      "age": 300,
      "species": "Celestial",
      "home_planet": "Alpha Centauri",
      "language": "Cosmic",
      "occupation": "Interstellar Diplomat"
    },
    {
      "name": "Astra",
      "age": 120,
      "species": "Stardust",
      "home_planet": "Polaris",
      "language": "Celestial",
      "occupation": "Astroarchaeologist"
    }
  ]
}
`

	var result map[string]any
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return result["aliens"].([]any)

}
