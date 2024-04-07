package models

import (
	"math/rand"
	"time"

	"github.com/andrewarrow/feedback/util"
	"github.com/brianvoe/gofakeit/v6"
)

type Field struct {
	Name      string   `json:"name"`
	Flavor    string   `json:"flavor"`
	Index     string   `json:"index"`
	Required  string   `json:"required"`
	Regex     string   `json:"regex"`
	Default   string   `json:"default"`
	Null      string   `json:"null"`
	JsonNames []string `json:"json_names"`
	JsonTypes []string `json:"json_types"`
}

func (f *Field) SqlTypeAndDefault() (string, string) {
	flavor := "varchar(255)"
	defaultString := "''"
	if f.Flavor == "int" {
		flavor = "int"
		defaultString = "0"
	} else if f.Flavor == "text" || f.Flavor == "json" || f.Flavor == "json_list" {
		flavor = "text"
	} else if f.Flavor == "bigint" {
		flavor = "bigint"
		defaultString = "0"
	} else if f.Flavor == "float" || f.Flavor == "double" {
		flavor = "DOUBLE PRECISION"
		defaultString = "0"
	} else if f.Flavor == "list" {
		flavor = "text"
	} else if f.Flavor == "uuid" || f.Flavor == "citext" {
		flavor = "citext"
	} else if f.Flavor == "enum" {
		defaultString = "'" + f.Default + "'"
	} else if f.Flavor == "geometry" {
		flavor = "geometry(Point,4326)"
	} else if f.Flavor == "bool" {
		flavor = "boolean"
		defaultString = "false"
	} else if f.Flavor == "timestamp" {
		flavor = "timestamp"
		defaultString = "NOW()"
	}
	if f.Null == "yes" {
		defaultString = "null"
	}
	return flavor, defaultString
}

func (f *Field) SaneDefault() any {
	if f.Flavor == "int" || f.Flavor == "bigint" {
		return 0
	} else if f.Flavor == "timestamp" && f.Null == "" {
		return time.Now()
	} else if f.Flavor == "bool" && f.Null == "" {
		return false
	} else if f.Null == "yes" {
		return nil
	}
	return ""
}

func (f *Field) CommonExclude() bool {
	if f.Name == "id" || f.Name == "created_at" {
		return true
	}
	if f.Name == "updated_at" || f.Name == "guid" {
		return true
	}
	return false
}

func (f *Field) RandomValue() any {
	var val any
	if f.Flavor == "uuid" {
		val = util.PseudoUuid()
	} else if f.Flavor == "username" {
		val = gofakeit.Username()
	} else if f.Flavor == "bool" {
		val = true
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
	} else if f.Flavor == "address_1_line" {
		val = gofakeit.Street() + " " + gofakeit.City() + ", " + gofakeit.StateAbr() + " " + gofakeit.Zip()
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
