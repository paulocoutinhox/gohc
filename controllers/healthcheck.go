package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/models/util"
	"log"
)

type HealthcheckController struct{}

func (This *HealthcheckController) Register() {
	app.Server.Router.GET("/healthcheck/list", This.HealthcheckList)
	log.Println("HealthcheckController register : OK")
}

func (This *HealthcheckController) HealthcheckList(c *gin.Context) {
	util.RenderTemplate(c.Writer, "healthcheck/list", nil)
}
