package utils

import (
	"html/template"

	"github.com/Dev-ManavSethi/my-website/models"
)

func ParseTemplates() error {

	models.Templates, models.DummyError = template.ParseGlob("../templates/*")
	if models.DummyError != nil {
		return models.DummyError
	}
	return nil

}
