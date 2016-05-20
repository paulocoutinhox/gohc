package main

import (
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/controllers"
	"github.com/prsolucoes/gohc/processor"
)

func main() {
	app.Server = app.NewWebServer()
	app.Server.LoadConfiguration()
	app.Server.LoadHealthchecks()
	app.Server.CreateBasicRoutes()

	{
		controller := controllers.HomeController{}
		controller.Register()
	}

	{
		controller := controllers.DashboardController{}
		controller.Register()
	}

	{
		controller := controllers.APIController{}
		controller.Register()
	}

	{
		controller := controllers.HealthcheckController{}
		controller.Register()
	}

	processor.CanRunHealthchecks = true
	processor.StartHealthcheckProcessor()

	app.Server.Start()
}
