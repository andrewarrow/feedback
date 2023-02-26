package controller

import "github.com/andrewarrow/feedback/models"

type Site struct {
	Phone  string         `json:"phone"`
	Models []models.Model `json:"models"`
}
