package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/models/util"
	"log"
)

type AboutController struct{}

func (This *AboutController) Register() {
	app.Server.Router.GET("/about", This.AboutIndex)
	log.Println("AboutController register : OK")
}

func (This *AboutController) AboutIndex(c *gin.Context) {
	util.RenderTemplate(c.Writer, "about/index", nil)
}
