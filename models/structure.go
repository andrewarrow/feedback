package models

import (
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/util"
	"github.com/brianvoe/gofakeit/v6"
)

const ISO8601 = "2006-01-02T15:04:05-07:00"
const HUMAN = "Monday, January 2, 2006 3:04 PM"

type Model struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
}

func (m *Model) TableName() string {
	prefix := os.Getenv("FEEDBACK_NAME")
	return prefix + "_" + util.Plural(m.Name)
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
	} else if f.Flavor == "timestamp" {
		flavor = "timestamp"
		defaultString = "NOW()"
	}
	return flavor, defaultString
}

func (f *Field) RandomValue() any {
	var val any
	if f.Flavor == "uuid" {
		val = util.PseudoUuid()
	} else if f.Flavor == "username" {
		val = gofakeit.Username()
	} else if f.Flavor == "name" {
		val = gofakeit.FirstName() + " " + gofakeit.LastName()
	} else if f.Flavor == "int" {
		val = rand.Intn(999)
	} else if f.Flavor == "timestamp" {
		dur := time.Duration(rand.Intn(24 * 7))
		if rand.Intn(2) == 0 {
			dur = dur * -1
		}
		val = time.Now().Add(time.Hour * dur).Format(ISO8601)
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
