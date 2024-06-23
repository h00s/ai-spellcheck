package initializers

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/tinylink/app/services"
)

func Services(c *raptor.Config) raptor.Services {
	return raptor.Services{
		services.NewGrammarService(),
	}
}
