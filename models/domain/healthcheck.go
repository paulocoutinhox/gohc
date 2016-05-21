package domain

import (
	"time"
	"log"
)

const (
	HEALTHCHECK_STATUS_SUCCESS = "success"
	HEALTHCHECK_STATUS_WARNING = "warning"
	HEALTHCHECK_STATUS_ERROR = "error"

	HEALTHCHECK_TYPE_PING = "ping"
	HEALTHCHECK_TYPE_RANGE = "range"
	HEALTHCHECK_TYPE_MANUAL = "manual"
)

type Healthcheck struct {
	Token            string    `json:"token"`
	Description      string    `json:"description"`
	LastUpdateAt     int64     `json:"lastUpdateAt"`
	Ping             int64     `json:"ping"`
	Range            float64   `json:"range"`
	Ranges           []float64 `json:"ranges"`
	Status           string    `json:"status"`
	Type             string    `json:"type"`
	WarningNotifiers []*HealthcheckNotifier `json:"warningNotifiers"`
	ErrorNotifiers   []*HealthcheckNotifier `json:"errorNotifiers"`
}

func (This *Healthcheck) Run() {
	if This.Type == HEALTHCHECK_TYPE_PING {
		This.UpdatePing()

		if This.InSuccessRange(float64(This.Ping)) {
			This.Status = HEALTHCHECK_STATUS_SUCCESS
		} else if This.InWarningRange(float64(This.Ping)) {
			This.Status = HEALTHCHECK_STATUS_WARNING
			This.NotifyWarningStatus()
		} else if This.InErrorRange(float64(This.Ping)) {
			This.Status = HEALTHCHECK_STATUS_ERROR
			This.NotifyErrorStatus()
		}
	} else if This.Type == HEALTHCHECK_TYPE_RANGE {
		if This.InSuccessRange(This.Range) {
			This.Status = HEALTHCHECK_STATUS_SUCCESS
		} else if This.InWarningRange(This.Range) {
			This.Status = HEALTHCHECK_STATUS_WARNING
			This.NotifyWarningStatus()
		} else if This.InErrorRange(This.Range) {
			This.Status = HEALTHCHECK_STATUS_ERROR
			This.NotifyErrorStatus()
		}
	} else if This.Type == HEALTHCHECK_TYPE_MANUAL {
		if This.Status == HEALTHCHECK_STATUS_WARNING {
			This.NotifyWarningStatus()
		} else if This.Status == HEALTHCHECK_STATUS_ERROR {
			This.NotifyErrorStatus()
		}
	}
}

func (This *Healthcheck) NotifyWarningStatus() {
	if This.WarningNotifiers != nil {
		for _, notifier := range This.WarningNotifiers {
			if notifier.CanSendNotification() {
				log.Println("Healthcheck : Started process to send warning notifications")
				// TODO: SEND ONLY TO THIS NOTIFIER CONFIG
				go NotifierManagerProcess(This)
			}
		}
	}
}

func (This *Healthcheck) NotifyErrorStatus() {
	if This.ErrorNotifiers != nil {
		for _, notifier := range This.ErrorNotifiers {
			if notifier.CanSendNotification() {
				log.Println("Healthcheck : Started process to send error notifications")
				// TODO: SEND ONLY TO THIS NOTIFIER CONFIG
				go NotifierManagerProcess(This)
			}
		}
	}
}

func (This *Healthcheck) SetLastUpdateAtCurrentTime() {
	This.LastUpdateAt = This.GetCurrentTimeInMS()
}

func (This *Healthcheck) UpdateLastPingData() {
	currentTime := This.GetCurrentTimeInMS()
	lastPingTime := This.LastUpdateAt
	This.Ping = currentTime - lastPingTime
	This.LastUpdateAt = currentTime
}

func (This *Healthcheck) UpdateLastRangeData(newRange float64) {
	currentTime := This.GetCurrentTimeInMS()
	This.Range = newRange
	This.LastUpdateAt = currentTime
}

func (This *Healthcheck) UpdatePing() {
	currentTime := This.GetCurrentTimeInMS()
	lastPingTime := This.LastUpdateAt
	This.Ping = currentTime - lastPingTime
}

func (This *Healthcheck) InSuccessRange(value float64) bool {
	return (value <= This.Ranges[0])
}

func (This *Healthcheck) InWarningRange(value float64) bool {
	if value <= This.Ranges[0] {
		return false
	}

	return (value <= This.Ranges[1])
}

func (This *Healthcheck) InErrorRange(value float64) bool {
	if value <= This.Ranges[0] {
		return false
	}

	return (value > This.Ranges[1])
}

func (This *Healthcheck) SetStatusSuccess() {
	This.Status = HEALTHCHECK_STATUS_SUCCESS
}

func (This *Healthcheck) SetStatusWarning() {
	This.Status = HEALTHCHECK_STATUS_WARNING
}

func (This *Healthcheck) SetStatusError() {
	This.Status = HEALTHCHECK_STATUS_ERROR
}

func (This *Healthcheck) GetCurrentTimeInMS() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Millisecond)
}
