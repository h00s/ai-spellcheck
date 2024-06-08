package config

import "github.com/go-raptor/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1",
			raptor.Scope("/grammar",
				raptor.Route("GET", "/check", "GrammarController", "Check"),
			),
		),
	)
}
