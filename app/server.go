package app

import (
	"encoding/json"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/prsolucoes/gohc/models/domain"
	"github.com/prsolucoes/gohc/processor"
	"io/ioutil"
	"log"
	"strconv"
)

type WebServer struct {
	Router       *gin.Engine
	Config       *ini.File
	Host         string
	WorkspaceDir string
	ResourcesDir string
}

var (
	Server *WebServer
)

func NewWebServer() *WebServer {
	server := new(WebServer)

	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()
	server.Router.Use(gin.Recovery())

	return server
}

func (This *WebServer) CreateBasicRoutes() {
	This.Router.Static("/static", This.ResourcesDir+"/static")
	log.Println("Router creation : OK")
}

func (This *WebServer) LoadHealthchecks() {
	fileName := This.WorkspaceDir + "/healthchecks.json"

	log.Printf("Loading healthcheck list file: %v", fileName)
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalf("Failed to load healthchecks file: %v", err)
	}

	healthcheckFile := domain.HealthchecksFile{}

	err = json.Unmarshal(file, &healthcheckFile)

	if err != nil {
		log.Fatalf("Failed to read healthchecks file: %v", err)
	}

	for i, healthcheck := range healthcheckFile.Healthchecks {
		if healthcheck.Type == domain.HEALTHCHECK_TYPE_PING {
			if healthcheck.Ranges == nil || len(healthcheck.Ranges) != 2 {
				log.Fatalf("Healthcheck (Token: %v, Index: %v) don't have 2 ranges", healthcheck.Token, i)
			}

			healthcheck.SetStatusSuccess()
			healthcheck.SetLastUpdateAtCurrentTime()
			healthcheck.UpdateLastPingData()
		} else if healthcheck.Type == domain.HEALTHCHECK_TYPE_RANGE {
			if healthcheck.Ranges == nil || len(healthcheck.Ranges) != 2 {
				log.Fatalf("Healthcheck (Token: %v, Index: %v) don't have 2 ranges", healthcheck.Token, i)
			}

			healthcheck.SetStatusSuccess()
			healthcheck.SetLastUpdateAtCurrentTime()
			healthcheck.UpdateLastRangeData(0)
		} else if healthcheck.Type == domain.HEALTHCHECK_TYPE_MANUAL {
			healthcheck.SetStatusSuccess()
			healthcheck.SetLastUpdateAtCurrentTime()
		} else {
			log.Fatalf("Healthcheck (Token: %v, Index: %v) has invalid type", healthcheck.Token, i)
		}
	}

	processor.Healthchecks = healthcheckFile.Healthchecks

	for i, notifier := range healthcheckFile.Notifiers {
		if notifier.Plugin == domain.NOTIFIER_PLUGIN_CLI_NAME {
			plugin := &domain.NotifierPluginCLI{
				ID:     notifier.ID,
				Params: notifier.Params,
			}

			domain.NotifierManagerAddPlugin(plugin)

			log.Printf("Notifier plugin (Id: %v, Index: %v) was added", notifier.ID, i)
		} else if notifier.Plugin == domain.NOTIFIER_PLUGIN_HTTP_GET_NAME {
			plugin := &domain.NotifierPluginHttpGet{
				ID:     notifier.ID,
				Params: notifier.Params,
			}

			domain.NotifierManagerAddPlugin(plugin)

			log.Printf("Notifier plugin (Id: %v, Index: %v) was added", notifier.ID, i)
		} else if notifier.Plugin == domain.NOTIFIER_PLUGIN_SENDGRID_NAME {
			plugin := &domain.NotifierPluginSendGrid{
				ID:     notifier.ID,
				Params: notifier.Params,
			}

			domain.NotifierManagerAddPlugin(plugin)

			log.Printf("Notifier plugin (Id: %v, Index: %v) was added", notifier.ID, i)
		} else if notifier.Plugin == domain.NOTIFIER_PLUGIN_PUSHBULLET_NAME {
			plugin := &domain.NotifierPluginPushBullet{
				ID:     notifier.ID,
				Params: notifier.Params,
			}

			domain.NotifierManagerAddPlugin(plugin)

			log.Printf("Notifier plugin (Id: %v, Index: %v) was added", notifier.ID, i)
		} else {
			log.Printf("Notifier plugin (Id: %v, Index: %v) is unknown", notifier.ID, i)
		}
	}

	log.Printf("Healthchecks file (%v) loaded", fileName)
}

func (This *WebServer) LoadConfiguration() {
	var configFileName = ""
	flag.StringVar(&configFileName, "f", "config.ini", "set config.ini location")
	flag.Parse()

	config, err := ini.Load([]byte(""), configFileName)

	if err == nil {
		This.Config = config

		serverSection, err := config.GetSection("server")

		if err != nil {
			This.Host = ":8080"
			This.WorkspaceDir = ""
			This.ResourcesDir = ""
			processor.WarmTime = 10
		} else {
			{
				// host
				host := serverSection.Key("host").Value()

				if host == "" {
					host = ":8080"
				}

				This.Host = host
			}

			{
				// workspace
				workspaceDir := serverSection.Key("workspaceDir").Value()
				This.WorkspaceDir = workspaceDir
			}

			{
				// resources dir
				resourcesDir := serverSection.Key("resourcesDir").Value()
				This.ResourcesDir = resourcesDir
			}

			{
				// warm time
				warmTime := serverSection.Key("warmTime").Value()
				value, err := strconv.ParseInt(warmTime, 10, 64)

				if err != nil {
					log.Fatalf("Configuration file load error : %s", err.Error())
				}

				processor.WarmTime = value
			}
		}

		log.Println("Configuration file load : OK")
	} else {
		log.Fatalf("Configuration file load error : %s", err.Error())
	}
}

func (This *WebServer) Start() {
	log.Printf("Open GoHC on your browser: %v", This.Host)
	err := This.Router.Run(This.Host)

	if err != nil {
		log.Fatalf("Server not started: %v", err)
	}
}
