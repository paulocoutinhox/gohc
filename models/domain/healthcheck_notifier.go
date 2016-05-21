package domain

import "time"

type HealthcheckNotifier struct {
	ID                 string                 `json:"id"`
	Interval           int64                  `json:"interval" default:"86400"`
	LastNotificationAt int64                  `json:"lastNotificationAt"`
	Params             map[string]interface{} `json:"params"`
}

func (This *HealthcheckNotifier) CanSendNotification() bool {
	currentTime := This.GetCurrentTimeInMS()

	if currentTime > (This.LastNotificationAt + This.Interval) {
		This.LastNotificationAt = currentTime
		return true
	}

	return false
}

func (This *HealthcheckNotifier) GetCurrentTimeInMS() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Millisecond)
}
