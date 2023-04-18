package router

import "github.com/andrewarrow/feedback/models"

type FeedbackSite struct {
	Footer string          `json:"footer"`
	Title  string          `json:"title"`
	Models []*models.Model `json:"models"`
	Routes []*models.Route `json:"routes"`
}

func (s *FeedbackSite) FindModel(id string) *models.Model {
	for _, m := range s.Models {
		if m.Name == id {
			return m
		}
	}

	return nil
}

func (s *FeedbackSite) FindField(model *models.Model, id string) *models.Field {
	for _, f := range model.Fields {
		if f.Name == id {
			return f
		}
	}

	return nil
}
