package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/processor"
	"github.com/prsolucoes/gowebresponse"
	"log"
	"strings"
)

type APIController struct{}

func (This *APIController) Register() {
	app.Server.Router.GET("/api/ping/:token", This.APIPing)
	app.Server.Router.GET("/api/healthcheck/count", This.APIHealthcheckCount)
	app.Server.Router.GET("/api/healthcheck/list", This.APIHealthcheckList)
	log.Println("APIController register : OK")
}

func (This *APIController) APIPing(c *gin.Context) {
	healthcheckToken := strings.Trim(c.Param("token"), "")
	healthcheck, err := processor.HealthcheckByToken(healthcheckToken)

	response := new(gowebresponse.WebResponse)

	if err == nil {
		healthcheck.UpdateLastPingData()

		response.Success = true
		response.Message = ""
	} else {
		response.Success = false
		response.Message = "error"
		response.AddDataError("error", err.Error())
	}

	c.JSON(200, response)
}

func (This *APIController) APIHealthcheckCount(c *gin.Context) {
	response := new(gowebresponse.WebResponse)
	response.Success = true
	response.Message = ""
	response.AddData("count", len(processor.Healthchecks))
	c.JSON(200, response)
}

func (This *APIController) APIHealthcheckList(c *gin.Context) {
	response := new(gowebresponse.WebResponse)
	response.Success = true
	response.Message = ""
	response.AddData("list", processor.Healthchecks)
	c.JSON(200, response)
}