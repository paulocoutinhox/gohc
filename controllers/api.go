package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/models/domain"
	"github.com/prsolucoes/gohc/processor"
	"github.com/prsolucoes/gowebresponse"
	"log"
	"strconv"
	"strings"
)

type APIController struct{}

func (This *APIController) Register() {
	app.Server.Router.GET("/api/update/ping/:token", This.APIUpdatePing)
	app.Server.Router.GET("/api/update/range/:token/:range", This.APIUpdateRange)
	app.Server.Router.GET("/api/update/manual/:token/:status", This.APIUpdateManual)
	app.Server.Router.GET("/api/healthcheck/count", This.APIHealthcheckCount)
	app.Server.Router.GET("/api/healthcheck/list", This.APIHealthcheckList)
	app.Server.Router.GET("/api/system/reload", This.APISystemReload)
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
	healthcheckRange := strings.Trim(c.Param("range"), "")
	healthcheck, err := processor.HealthcheckByToken(healthcheckToken)

	response := new(gowebresponse.WebResponse)

	if err == nil {
		if healthcheck.Type == domain.HEALTHCHECK_TYPE_RANGE {
			newRange, err := strconv.ParseFloat(healthcheckRange, 64)

			if err == nil {
				healthcheck.UpdateLastRangeData(newRange)

				response.Success = true
				response.Message = "updated"
			} else {
				response.Success = false
				response.Message = "error"
				response.AddDataError("error", "Invalid range value")
			}
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

func (This *APIController) APIUpdateManual(c *gin.Context) {
	healthcheckToken := strings.Trim(c.Param("token"), "")
	healthcheckStatus := strings.Trim(c.Param("status"), "")
	healthcheck, err := processor.HealthcheckByToken(healthcheckToken)

	response := new(gowebresponse.WebResponse)

	if err == nil {
		if healthcheck.Type == domain.HEALTHCHECK_TYPE_MANUAL {
			if healthcheckStatus == domain.HEALTHCHECK_STATUS_SUCCESS {
				healthcheck.SetStatusSuccess()

				response.Success = true
				response.Message = "updated"
			} else if healthcheckStatus == domain.HEALTHCHECK_STATUS_WARNING {
				healthcheck.SetStatusWarning()

				response.Success = true
				response.Message = "updated"
			} else if healthcheckStatus == domain.HEALTHCHECK_STATUS_ERROR {
				healthcheck.SetStatusError()

				response.Success = true
				response.Message = "updated"
			} else {
				response.Success = false
				response.Message = "error"
				response.AddDataError("error", "Invalid status")
			}
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
	response.AddData("count", len(processor.Healthchecks))
	c.JSON(200, response)
}

func (This *APIController) APISystemReload(c *gin.Context) {
	err := app.Server.TestHealthchecksFile(true)
	response := new(gowebresponse.WebResponse)

	if err == nil {
		response.Success = true
		response.Message = ""
	} else {
		response.Success = false
		response.Message = "error"
		response.AddDataError("error", err.Error())
	}

	c.JSON(200, response)
}
