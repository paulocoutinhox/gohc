package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Configuration struct {
	Server       *ConfigurationServer `json:"server"`
	Healthchecks []*Healthcheck       `json:"healthchecks"`
	Notifiers    []*Notifier          `json:"notifiers"`
}

type ConfigurationServer struct {
	Host     string `json:"host"`
	WarmTime int64  `json:"warmTime"`
}

func NewConfigurationFromFile(file string) (*Configuration, error) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to load configuration file: %v", err))
	}

	configuration := Configuration{}

	err = json.Unmarshal(data, &configuration)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to read configuration file: %v", err))
	}

	return &configuration, nil
}
