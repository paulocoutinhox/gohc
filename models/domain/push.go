package domain

import (
	"strconv"
	"strings"
)

func PushCreateTitle(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) string {
	title := "New healthcheck [status] alert!"
	title = strings.Replace(title, "[status]", healthcheck.Status, -1)
	return title
}

func PushCreateMessage(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) string {
	body := ""

	if healthcheck.Type == HEALTHCHECK_TYPE_MANUAL {
		body = `New [status] alert! Healthcheck: "[description]", Type: "[type-text]", Status: "[status]"`
	} else {
		body = `New [status] alert! Healthcheck: "[description]", Type: "[type-text]", Value: [type-value], Status: "[status]"`
	}

	body = strings.Replace(body, "[status]", healthcheck.Status, -1)
	body = strings.Replace(body, "[description]", healthcheck.Description, -1)

	if healthcheck.Type == HEALTHCHECK_TYPE_PING {
		body = strings.Replace(body, "[type-text]", "Ping", -1)
		body = strings.Replace(body, "[type-value]", strconv.FormatInt(healthcheck.Ping, 10), -1)
	} else if healthcheck.Type == HEALTHCHECK_TYPE_RANGE {
		body = strings.Replace(body, "[type-text]", "Range", -1)
		body = strings.Replace(body, "[type-value]", strconv.FormatFloat(healthcheck.Range, 'f', 2, 64), -1)
	}

	if healthcheck.Status == HEALTHCHECK_STATUS_SUCCESS {
		body = strings.Replace(body, "[status-color]", "#1ab394", -1)
	} else if healthcheck.Status == HEALTHCHECK_STATUS_WARNING {
		body = strings.Replace(body, "[status-color]", "#f8ac59", -1)
	} else if healthcheck.Status == HEALTHCHECK_STATUS_ERROR {
		body = strings.Replace(body, "[status-color]", "#ed5565", -1)
	}

	return body
}
