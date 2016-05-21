package domain

import (
	"io/ioutil"
	"log"
	"net/http"
)

const (
	NOTIFIER_PLUGIN_HTTP_GET_NAME = "httpget"
)

type NotifierPluginHttpGet struct {
	Params map[string]interface{}
	ID     string
}

func (This *NotifierPluginHttpGet) Notify(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) {
	log.Println("NotifierPluginHttpGet : Notify")

	if This.Params == nil {
		log.Println("NotifierPluginHttpGet : cannot execute because params is empty")
		return
	}

	url := ""

	for key, value := range This.Params {
		if key == "url" {
			if data, ok := value.(string); ok {
				url = data
			}
		}
	}

	log.Printf("NotifierPluginHttpGet : http get request to be execute (url: %v)", url)
	content, err := This.ExecuteHttpGet(url)

	if err != nil {
		log.Printf("NotifierPluginHttpGet : executed with error: %v", err)
	} else {
		log.Printf("NotifierPluginHttpGet : executed with output: \n %v", string(content))
	}
}

func (This *NotifierPluginHttpGet) GetName() string {
	return NOTIFIER_PLUGIN_HTTP_GET_NAME
}

func (This *NotifierPluginHttpGet) GetId() string {
	return This.ID
}

func (This *NotifierPluginHttpGet) ExecuteHttpGet(url string) (string, error) {
	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return string(content), nil
}
