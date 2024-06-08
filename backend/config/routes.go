package config

import "github.com/go-raptor/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1",
			raptor.Scope("/spellcheck",
				raptor.Route("GET", "/hello", "SpellcheckController", "Hello"),
			),
		),
	)
}
