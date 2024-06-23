package config

import "github.com/go-raptor/raptor/v2"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1",
			raptor.Scope("/grammar",
				raptor.Route("POST", "/check", "GrammarController", "Check"),
			),
		),
	)
}
