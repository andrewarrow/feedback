package router

import "github.com/andrewarrow/feedback/models"

type FeedbackSite struct {
	Footer  string          `json:"footer"`
	Title   string          `json:"title"`
	Models  []*models.Model `json:"models"`
	Routes  []*models.Route `json:"routes"`
	Dynamic []*models.Model `json:"dynamic"`
}

func (s *FeedbackSite) FindModel(id string) *models.Model {
	for _, m := range s.Models {
		if m.Name == id {
			return m
		}
	}

	return nil
}

func (s *FeedbackSite) FindModelOrDynamic(id string) *models.Model {
	m := s.FindModel(id)
	if m == nil {
		m = s.FindDynamic(id)
	}
	return m
}

func (s *FeedbackSite) FindDynamic(id string) *models.Model {
	for _, m := range s.Dynamic {
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
