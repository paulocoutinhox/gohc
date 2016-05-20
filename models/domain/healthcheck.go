package domain

import (
	"time"
)

const (
	HEALTHCHECK_STATUS_SUCCESS = "success"
	HEALTHCHECK_STATUS_WARNING = "warning"
	HEALTHCHECK_STATUS_ERROR   = "error"

	HEALTHCHECK_TYPE_PING   = "ping"
	HEALTHCHECK_TYPE_RANGE  = "range"
	HEALTHCHECK_TYPE_MANUAL = "manual"
)

type Healthcheck struct {
	Token       string    `json:"token"`
	Description string    `json:"description"`
	LastPingAt  int64     `json:"lastPingAt"`
	LastRangeAt int64     `json:"lastRangeAt"`
	Ping        int64     `json:"ping"`
	Range       float64   `json:"range"`
	Ranges      []float64 `json:"ranges"`
	Status      string    `json:"status"`
	Type        string    `json:"type"`
}

func (This *Healthcheck) Run() {
	if This.Type == HEALTHCHECK_TYPE_PING {
		This.UpdatePing()

		if This.InSuccessRange(float64(This.Ping)) {
			This.Status = HEALTHCHECK_STATUS_SUCCESS
		} else if This.InWarningRange(float64(This.Ping)) {
			This.Status = HEALTHCHECK_STATUS_WARNING
		} else if This.InErrorRange(float64(This.Ping)) {
			This.Status = HEALTHCHECK_STATUS_ERROR
		}
	} else if This.Type == HEALTHCHECK_TYPE_RANGE {
		if This.InSuccessRange(This.Range) {
			This.Status = HEALTHCHECK_STATUS_SUCCESS
		} else if This.InWarningRange(This.Range) {
			This.Status = HEALTHCHECK_STATUS_WARNING
		} else if This.InErrorRange(This.Range) {
			This.Status = HEALTHCHECK_STATUS_ERROR
		}
	}
}

func (This *Healthcheck) UpdateLastPingData() {
	currentTime := This.GetCurrentTimeInMS()
	lastPingTime := This.LastPingAt
	This.Ping = currentTime - lastPingTime
	This.LastPingAt = currentTime
}

func (This *Healthcheck) UpdateLastRangeData(newRange float64) {
	currentTime := This.GetCurrentTimeInMS()
	This.Range = newRange
	This.LastRangeAt = currentTime
}

func (This *Healthcheck) UpdatePing() {
	currentTime := This.GetCurrentTimeInMS()
	lastPingTime := This.LastPingAt
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
