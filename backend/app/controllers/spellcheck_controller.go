package controllers

import "github.com/go-raptor/raptor"

type SpellcheckController struct {
	raptor.Controller
}

func (sc *SpellcheckController) Hello(c *raptor.Context) error {
	return c.JSON(raptor.Map{
		"message": "Hello, World!",
	})
}
