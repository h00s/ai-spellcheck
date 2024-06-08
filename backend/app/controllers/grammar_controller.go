package controllers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
)

type GrammarController struct {
	raptor.Controller
}

func (sc *GrammarController) Check(c *raptor.Context) error {
	var link models.Content
	if err := c.BodyParser(&link); err != nil {
		return c.JSONError(raptor.NewErrorBadRequest("Invalid JSON"))
	}
	return c.JSON(link)
}
