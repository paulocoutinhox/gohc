package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/models/util"
	"log"
)

type DashboardController struct{}

func (This *DashboardController) Register() {
	app.Server.Router.GET("/dashboard", This.HomeIndex)
	log.Println("DashboardController register : OK")
}

func (This *DashboardController) HomeIndex(c *gin.Context) {
	util.RenderTemplate(c.Writer, "dashboard/index", nil)
}
