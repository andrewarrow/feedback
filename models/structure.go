package models

import (
	"github.com/andrewarrow/feedback/prefix"
	"github.com/andrewarrow/feedback/util"
)

const ISO8601 = "2006-01-02T15:04:05-07:00"
const HUMAN = "Monday, January 2, 2006 3:04 PM"

type Model struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
}

func (m *Model) EnsureIdAndCreatedAt() {
	ca := FindField(m, "created_at")
	if ca == nil {
		f := Field{}
		f.Name = "created_at"
		f.Flavor = "timestamp"
		f.Index = "yes"
		m.Fields = append(m.Fields, &f)
	}
	ua := FindField(m, "updated_at")
	if ua == nil {
		f := Field{}
		f.Name = "updated_at"
		f.Flavor = "timestamp"
		f.Index = "yes"
		m.Fields = append(m.Fields, &f)
	}
	id := FindField(m, "id")
	if id == nil {
		f := Field{}
		f.Name = "id"
		f.Flavor = "int"
		m.Fields = append(m.Fields, &f)
	}
	guid := FindField(m, "guid")
	if guid == nil {
		f := Field{}
		f.Name = "guid"
		f.Flavor = "uuid"
		f.Index = "yes"
		m.Fields = append(m.Fields, &f)
	}
}

func (m *Model) TableName() string {
	return prefix.Tablename(util.Plural(m.Name))
}

func TypeToFlavor(dt string) string {
	if dt == "bigint" || dt == "integer" {
		return "int"
	} else if dt == "boolean" {
		return "bool"
	} else if dt == "text" {
		return "text"
	}
	return "oneWord"
}
