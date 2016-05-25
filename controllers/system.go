package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/models/util"
	"log"
)

type SystemController struct{}

func (This *SystemController) Register() {
	app.Server.Router.GET("/system/settings", This.SystemSettings)
	log.Println("SystemController register : OK")
}

func (This *SystemController) SystemSettings(c *gin.Context) {
	util.RenderTemplate(c.Writer, "system/settings", nil)
}
