package models

type Model struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
	Index  string `json:"index"`
}

func (f *Field) SqlTypeAndDefault() (string, string) {
	flavor := "varchar(255)"
	defaultString := "''"
	if f.Flavor == "int" {
		flavor = "int"
		defaultString = "0"
	} else if f.Flavor == "text" {
		flavor = "text"
	}
	return flavor, defaultString
}
