package router

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Building struct {
	Timestamp string
	Ago       string
	Guid      string
	Name      string
	Address   string
	City      string
	State     string
	Postal    string
	Phone     string
	Url       string
	Country   string
	About     string
	Units     int64
}

func FetchBuildings(db *sqlx.DB) []*Building {
	items := []*Building{}
	rows, err := db.Queryx("SELECT * FROM buildings ORDER BY created_at desc limit 30")
	if err != nil {
		return items
	}
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		b := buildingFromMap(m)
		items = append(items, b)
	}
	return items
}

func buildingFromMap(m map[string]any) *Building {
	building := Building{}
	building.Url = fmt.Sprintf("%s", m["url"])
	building.Guid = fmt.Sprintf("%s", m["guid"])
	building.Name = fmt.Sprintf("%s", m["name"])
	building.Units = m["units"].(int64)
	building.About = fmt.Sprintf("%s", m["about"])
	building.Country = fmt.Sprintf("%s", m["country"])
	building.City = fmt.Sprintf("%s", m["city"])
	building.State = fmt.Sprintf("%s", m["state"])
	building.Address = fmt.Sprintf("%s", m["address"])
	building.Url = fmt.Sprintf("%s", m["url"])

	building.Timestamp, building.Ago = FixTime(m)
	return &building
}
