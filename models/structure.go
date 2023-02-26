package models

type Model struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
}
