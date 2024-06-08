package controllers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/app/services"
)

type GrammarController struct {
	raptor.Controller

	Grammar *services.GrammarService
}

func (sc *GrammarController) Check(c *raptor.Context) error {
	var link models.Content
	if err := c.BodyParser(&link); err != nil {
		return c.JSONError(raptor.NewErrorBadRequest("Invalid JSON"))
	}

	validatedContent, err := sc.Grammar.Check(link)
	if err != nil {
		return c.JSONError(err)
	}
	return c.JSON(validatedContent)
}
