package domain

import "log"

var (
	Plugins []INotifierPlugin
)

func NotifierManagerAddPlugin(plugin INotifierPlugin) {
	Plugins = append(Plugins, plugin)
}

func NotifierManagerProcess(healthcheck *Healthcheck) error {
	log.Println("NotifierPluginManager : NotifierManagerProcess")

	healthcheckNotifiers := []*HealthcheckNotifier{}

	if healthcheck.Status == HEALTHCHECK_STATUS_WARNING {
		healthcheckNotifiers = healthcheck.WarningNotifiers
	} else if healthcheck.Status == HEALTHCHECK_STATUS_ERROR {
		healthcheckNotifiers = healthcheck.ErrorNotifiers
	}

	if healthcheckNotifiers == nil {
		return nil
	}

	for _, plugin := range Plugins {
		for _, healthcheckNotifier := range healthcheckNotifiers {
			if healthcheckNotifier.ID == plugin.GetId() {
				plugin.Notify(*healthcheck)
			}
		}
	}

	return nil
}
