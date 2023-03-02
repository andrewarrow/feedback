package router

import "github.com/andrewarrow/feedback/models"

type Site struct {
	Phone  string         `json:"phone"`
	Models []models.Model `json:"models"`
}

func (s *Site) FindModel(id string) *models.Model {
	for _, m := range s.Models {
		if m.Name == id {
			return &m
		}
	}

	return nil
}
