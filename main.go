package main

import (
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/controllers"
	"github.com/prsolucoes/gohc/processor"
)

func main() {
	app.Server = app.NewWebServer()
	app.Server.LoadConfiguration()
	app.Server.TestHealthchecksFile(true)
	app.Server.CreateBasicRoutes()

	{
		controller := controllers.APIController{}
		controller.Register()
	}

	processor.CanRunHealthchecks = true
	processor.StartHealthcheckProcessor()

	app.Server.Start()
}
