package models

import (
	"strings"

	"github.com/andrewarrow/feedback/prefix"
	"github.com/andrewarrow/feedback/util"
)

const ISO8601 = "2006-01-02T15:04:05-07:00"
const HUMAN = "Monday, January 2, 2006 3:04 PM"
const HUMAN_DATE = "2006-01-02"

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

func TypeToFlavor(dt, udt, cd string) string {
	if dt == "bigint" || dt == "integer" {
		return "int"
	} else if dt == "boolean" {
		return "bool"
	} else if dt == "text" {
		return "text"
	} else if dt == "real" {
		return "double"
	} else if dt == "numeric" {
		return "numeric"
	} else if udt == "timestamptz" {
		return "timestamp"
	} else if udt == "geometry" {
		return "geometry"
	} else if dt == "jsonb" && cd == "" {
		return "json"
	} else if dt == "jsonb" && strings.Contains(cd, "[]") {
		return "json_list"
	} else if dt == "jsonb" && strings.Contains(cd, "{}") {
		return "json"
	} else if strings.HasPrefix(udt, "enum_") {
		return udt
	}
	return "name"
}
