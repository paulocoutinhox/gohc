package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/models/domain"
	"github.com/prsolucoes/gohc/processor"
	"github.com/prsolucoes/gowebresponse"
	"log"
	"strings"
)

type APIController struct{}

func (This *APIController) Register() {
	app.Server.Router.GET("/api/update/ping/:token", This.APIUpdatePing)
	app.Server.Router.GET("/api/healthcheck/count", This.APIHealthcheckCount)
	app.Server.Router.GET("/api/healthcheck/list", This.APIHealthcheckList)
	log.Println("APIController register : OK")
}

func (This *APIController) APIUpdatePing(c *gin.Context) {
	healthcheckToken := strings.Trim(c.Param("token"), "")
	healthcheck, err := processor.HealthcheckByToken(healthcheckToken)

	response := new(gowebresponse.WebResponse)

	if err == nil {
		if healthcheck.Type == domain.HEALTHCHECK_TYPE_PING {
			healthcheck.UpdateLastPingData()

			response.Success = true
			response.Message = "updated"
		} else {
			response.Success = false
			response.Message = "error"
			response.AddDataError("error", "Invalid type")
		}
	} else {
		response.Success = false
		response.Message = "error"
		response.AddDataError("error", err.Error())
	}

	c.JSON(200, response)
}

func (This *APIController) APIUpdateRange(c *gin.Context) {
	healthcheckToken := strings.Trim(c.Param("token"), "")
	healthcheck, err := processor.HealthcheckByToken(healthcheckToken)

	response := new(gowebresponse.WebResponse)

	if err == nil {
		if healthcheck.Type == domain.HEALTHCHECK_TYPE_RANGE {
			healthcheck.UpdateLastPingData()

			response.Success = true
			response.Message = "updated"
		} else {
			response.Success = false
			response.Message = "error"
			response.AddDataError("error", "Invalid type")
		}
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
