package models

import (
	"fmt"
	"math/rand"

	"github.com/andrewarrow/feedback/util"
	"github.com/brianvoe/gofakeit/v6"
)

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

func (f *Field) RandomValue() any {
	var val any
	if f.Flavor == "uuid" {
		val = util.PseudoUuid()
	} else if f.Flavor == "username" {
		val = gofakeit.Username()
	} else if f.Flavor == "int" {
		val = fmt.Sprintf("%d", rand.Intn(999))
	} else if f.Flavor == "int" {
		val = fmt.Sprintf("%d", rand.Intn(999))
	} else {
		val = gofakeit.Word()
	}

	return val
}
