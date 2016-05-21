package processor

import (
	"errors"
	"github.com/prsolucoes/gohc/models/domain"
	"log"
	"time"
)

var (
	Healthchecks                []*domain.Healthcheck
	CanRunHealthchecks          bool
	HealthchecksProcessorTicker *time.Ticker
	StartedAt                   int64
	WarmTime                    int64
)

func StartHealthcheckProcessor() {
	StartedAt = time.Now().Unix()
	HealthchecksProcessorTicker = time.NewTicker(time.Second * 1)

	go func() {
		for range HealthchecksProcessorTicker.C {
			if CanRunHealthchecks {
				if OutOfWarmTime() {
					for _, healthcheck := range Healthchecks {
						go healthcheck.Run()
					}
				}
			}
		}
	}()

	log.Printf("Healthcheck processor started : OK")
}

func OutOfWarmTime() bool {
	currentTime := time.Now().Unix()

	if (currentTime - StartedAt) > WarmTime {
		return true
	}

	return false
}

func HealthcheckByToken(token string) (*domain.Healthcheck, error) {
	for _, healthcheck := range Healthchecks {
		if healthcheck.Token == token {
			return healthcheck, nil
		}
	}

	return nil, errors.New("Healthcheck not found by this token")
}
