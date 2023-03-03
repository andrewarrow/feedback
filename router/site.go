package router

import "github.com/andrewarrow/feedback/models"

type Site struct {
	Phone  string          `json:"phone"`
	Models []*models.Model `json:"models"`
}

func (s *Site) FindModel(id string) *models.Model {
	for _, m := range s.Models {
		if m.Name == id {
			return m
		}
	}

	return nil
}

func (s *Site) AddField(id string, field models.Field) {
	for _, m := range s.Models {
		if m.Name == id {
			m.Fields = append(m.Fields, field)
			break
		}
	}
}
