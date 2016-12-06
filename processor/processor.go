package processor

import (
	"errors"
	"github.com/prsolucoes/gohc/app"
	"github.com/prsolucoes/gohc/models/domain"
	"github.com/prsolucoes/gohc/models/warm"
	"log"
	"time"
)

var (
	CanRunHealthchecks          bool
	HealthchecksProcessorTicker *time.Ticker
)

func StartHealthcheckProcessor() {
	warm.StartedAt = time.Now().UTC().UnixNano() / int64(time.Millisecond)
	HealthchecksProcessorTicker = time.NewTicker(time.Second * 1)

	go func() {
		for range HealthchecksProcessorTicker.C {
			if CanRunHealthchecks {
				for _, healthcheck := range app.Server.Configuration.Healthchecks {
					go healthcheck.Run()
				}
			}
		}
	}()

	log.Println("Healthcheck processor started : OK")
}

func HealthcheckByToken(token string) (*domain.Healthcheck, error) {
	for _, healthcheck := range app.Server.Configuration.Healthchecks {
		if healthcheck.Token == token {
			return healthcheck, nil
		}
	}

	return nil, errors.New("Healthcheck not found by this token")
}
