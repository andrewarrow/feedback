package models

import (
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
		val = rand.Intn(999)
	} else if f.Flavor == "oneWord" {
		val = gofakeit.Word()
	} else if f.Flavor == "fewWords" {
		val = gofakeit.Word() + " " + gofakeit.Word() + " " + gofakeit.Word()
	} else if f.Flavor == "oneWord" {
		val = gofakeit.Word()
	} else if f.Flavor == "address" {
		val = gofakeit.Street()
	} else if f.Flavor == "city" {
		val = gofakeit.City()
	} else if f.Flavor == "state" {
		val = gofakeit.StateAbr()
	} else if f.Flavor == "postal" {
		val = gofakeit.Zip()
	} else if f.Flavor == "country" {
		val = gofakeit.Country()
	} else if f.Flavor == "url" {
		val = gofakeit.URL()
	} else if f.Flavor == "firstName" {
		val = gofakeit.FirstName()
	} else if f.Flavor == "lastName" {
		val = gofakeit.LastName()
	} else if f.Flavor == "phone" {
		val = gofakeit.PhoneFormatted()
	} else if f.Flavor == "text" {
		val = gofakeit.LoremIpsumParagraph(1, 3, 33, ".")
	} else {
		val = gofakeit.Word()
	}
	return val
}
