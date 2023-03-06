package router

import "github.com/andrewarrow/feedback/models"

type FeedbackSite struct {
	Footer string          `json:"footer"`
	Title  string          `json:"title"`
	Models []*models.Model `json:"models"`
}

func (s *FeedbackSite) FindModel(id string) *models.Model {
	for _, m := range s.Models {
		if m.Name == id {
			return m
		}
	}

	return nil
}
