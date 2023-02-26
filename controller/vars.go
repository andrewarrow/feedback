package controller

import "html/template"

type Vars struct {
	Title  string
	Header template.HTML
	Footer template.HTML
	Phone  string
}

func NewVars() Vars {
	v := Vars{}
	v.Title = "Feedback"
	return v
}
