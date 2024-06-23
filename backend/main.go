package main

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/tinylink/config"
	"github.com/h00s/tinylink/config/initializers"
)

func main() {
	r := raptor.NewRaptor()

	r.Init(initializers.App(r.Utils.Config))
	r.RegisterRoutes(config.Routes())

	r.Listen()
}
