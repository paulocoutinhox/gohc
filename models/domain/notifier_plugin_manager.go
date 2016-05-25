package domain

import "log"

var (
	Plugins []INotifierPlugin
)

func NotifierManagerAddPlugin(plugin INotifierPlugin) {
	Plugins = append(Plugins, plugin)
}

func NotifierManagerClearPlugins() {
	Plugins = []INotifierPlugin{}
}

func NotifierManagerProcess(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) error {
	log.Println("NotifierPluginManager : NotifierManagerProcess")

	for _, plugin := range Plugins {
		if healthcheckNotifier.ID == plugin.GetId() {
			plugin.Notify(healthcheck, healthcheckNotifier)
		}
	}

	return nil
}
