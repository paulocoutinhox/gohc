package domain

import (
	"time"
)

const (
	CHECK_STATUS_SUCCESS = "success"
	CHECK_STATUS_WARNING = "warning"
	CHECK_STATUS_ERROR = "error"
)

type Healthcheck struct {
	Token       string        `json:"token"`
	Description string        `json:"description"`
	LastPingAt  int64         `json:"lastPingAt"`
	Ping        int64         `json:"ping"`
	Ranges      []int64       `json:"ranges"`
	Status      string        `json:"status"`
}

func (This *Healthcheck) Run() {
	This.UpdatePing()

	if This.InSuccessRange() {
		This.Status = CHECK_STATUS_SUCCESS
	} else if This.InWarningRange() {
		This.Status = CHECK_STATUS_WARNING
	} else if This.InErrorRange() {
		This.Status = CHECK_STATUS_ERROR
	}
}

func (This *Healthcheck) UpdateLastPingData() {
	currentTime := This.GetCurrentTimeInMS()
	lastPingTime := This.LastPingAt
	This.Ping = currentTime - lastPingTime
	This.LastPingAt = currentTime
}

func (This *Healthcheck) UpdatePing() {
	currentTime := This.GetCurrentTimeInMS()
	lastPingTime := This.LastPingAt
	This.Ping = currentTime - lastPingTime
}

func (This *Healthcheck) InSuccessRange() bool {
	return (This.Ping <= This.Ranges[0])
}

func (This *Healthcheck) InWarningRange() bool {
	if This.Ping <= This.Ranges[0] {
		return false
	}

	return (This.Ping <= This.Ranges[1])
}

func (This *Healthcheck) InErrorRange() bool {
	if This.Ping <= This.Ranges[0] {
		return false
	}

	return (This.Ping > This.Ranges[1])
}

func (This *Healthcheck) SetStatusSuccess() {
	This.Status = CHECK_STATUS_SUCCESS
}

func (This *Healthcheck) SetStatusWarning() {
	This.Status = CHECK_STATUS_WARNING
}

func (This *Healthcheck) SetStatusError() {
	This.Status = CHECK_STATUS_ERROR
}

func (This *Healthcheck) GetCurrentTimeInMS() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Millisecond)
}