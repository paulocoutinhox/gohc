package app

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/gohc/assets"
	"github.com/prsolucoes/gohc/models/domain"
	"github.com/prsolucoes/gohc/models/warm"
	"log"
)

type WebServer struct {
	Router            *gin.Engine
	Configuration     *domain.Configuration
	ConfigurationFile string
}

var (
	Server *WebServer
)

func NewWebServer() *WebServer {
	server := new(WebServer)

	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()
	server.Router.Use(gin.Recovery())
	server.Router.Use(gzip.Gzip(gzip.DefaultCompression))

	return server
}

func (This *WebServer) CreateBasicRoutes() {
	This.Router.NoRoute(This.RouteGeneral)
	This.Router.Use(static.Serve("/web-app", BinaryFileSystem("web-app")))
	log.Println("Router creation : OK")
}

func (This *WebServer) RouteGeneral(c *gin.Context) {
	data, err := assets.Asset("web-app/index.html")

	if err != nil {
		// asset was not found.
	}

	c.Data(200, "text/html", data)
}

func (This *WebServer) LoadHealthchecks(healthchecks []*domain.Healthcheck, notifiers []*domain.Notifier) error {
	if healthchecks == nil {
		return errors.New("Healthcheck list is invalid")
	}

	if notifiers == nil {
		return errors.New("Notifier list is invalid")
	}

	log.Printf("Loading %v healthchecks...", len(healthchecks))

	for i, healthcheck := range healthchecks {
		if healthcheck.Type == domain.HEALTHCHECK_TYPE_PING {
			healthcheck.SetStatusSuccess()
			healthcheck.SetLastUpdateAtCurrentTime()
			healthcheck.UpdateLastPingData()
			log.Printf("Healthcheck (Id: %v, Index: %v) was added", healthcheck.Token, i)
		} else if healthcheck.Type == domain.HEALTHCHECK_TYPE_RANGE {
			healthcheck.SetStatusSuccess()
			healthcheck.SetLastUpdateAtCurrentTime()
			healthcheck.UpdateLastRangeData(0)
			log.Printf("Healthcheck (Id: %v, Index: %v) was added", healthcheck.Token, i)
		} else if healthcheck.Type == domain.HEALTHCHECK_TYPE_MANUAL {
			healthcheck.SetStatusSuccess()
			healthcheck.SetLastUpdateAtCurrentTime()
			log.Printf("Healthcheck (Id: %v, Index: %v) was added", healthcheck.Token, i)
		}
	}

	log.Printf("Loading %v notifiers plugin...", len(notifiers))

	domain.NotifierManagerClearPlugins()

	for i, notifier := range notifiers {
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
		} else if notifier.Plugin == domain.NOTIFIER_PLUGIN_SLACK_WEBHOOK_NAME {
			plugin := &domain.NotifierPluginSlackWebHook{
				ID:     notifier.ID,
				Params: notifier.Params,
			}

			domain.NotifierManagerAddPlugin(plugin)

			log.Printf("Notifier plugin (Id: %v, Index: %v) was added", notifier.ID, i)
		} else {
			log.Printf("Notifier plugin (Id: %v, Index: %v) is unknown", notifier.ID, i)
		}
	}

	This.Configuration.Healthchecks = healthchecks
	This.Configuration.Notifiers = notifiers

	log.Println("Data was loaded with success!")

	return nil
}

func (This *WebServer) TestHealthchecksFile(load bool) error {
	log.Printf("Loading healthcheck list from file: %v...", This.ConfigurationFile)

	// read configuration file
	configuration, err := domain.NewConfigurationFromFile(This.ConfigurationFile)

	if err != nil {
		log.Printf("Failed to load configuration file: %v", err)
	}

	healthchecks := configuration.Healthchecks
	notifiers := configuration.Notifiers

	for i, healthcheck := range healthchecks {
		if healthcheck.Type == domain.HEALTHCHECK_TYPE_PING {
			if healthcheck.Ranges == nil || len(healthcheck.Ranges) != 2 {
				return errors.New(fmt.Sprintf("Healthcheck (Token: %v, Index: %v) don't have 2 ranges", healthcheck.Token, i))
			}
		} else if healthcheck.Type == domain.HEALTHCHECK_TYPE_RANGE {
			if healthcheck.Ranges == nil || len(healthcheck.Ranges) != 2 {
				return errors.New(fmt.Sprintf("Healthcheck (Token: %v, Index: %v) don't have 2 ranges", healthcheck.Token, i))
			}
		} else if healthcheck.Type == domain.HEALTHCHECK_TYPE_MANUAL {
			// ?
		} else {
			return errors.New(fmt.Sprintf("Healthcheck (Token: %v, Index: %v) has invalid type", healthcheck.Token, i))
		}
	}

	log.Printf("Healthchecks file (%v) tested : OK", This.ConfigurationFile)

	if load {
		return This.LoadHealthchecks(healthchecks, notifiers)
	}

	return nil
}

func (This *WebServer) LoadConfiguration() {
	// load from args
	var configFileName = ""
	flag.StringVar(&configFileName, "f", "config.json", "set config.json location")
	flag.Parse()

	if len(configFileName) > 0 {
		This.ConfigurationFile = configFileName
	} else {
		log.Fatal("Failed to load configuration file")
	}

	// read configuration file
	configuration, err := domain.NewConfigurationFromFile(This.ConfigurationFile)

	if err != nil {
		log.Fatalf("Failed to load configuration file: %v", err)
	}

	This.Configuration = configuration

	// warm time
	warm.WarmTime = This.Configuration.Server.WarmTime
}

func (This *WebServer) Start() {
	log.Printf("Open GoHC on your browser: %v", This.Configuration.Server.Host)
	err := This.Router.Run(This.Configuration.Server.Host)

	if err != nil {
		log.Fatalf("Server not started: %v", err)
	}
}
