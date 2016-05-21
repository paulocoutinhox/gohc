package domain

import (
	"github.com/bluele/slack"
	"log"
)

const (
	NOTIFIER_PLUGIN_SLACK_WEBHOOK_NAME = "slackwebhook"
)

type NotifierPluginSlackWebHook struct {
	Params map[string]interface{}
	ID     string
}

func (This *NotifierPluginSlackWebHook) Notify(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) {
	log.Println("NotifierPluginSlackWebHook : Notify")

	if This.Params == nil {
		log.Println("NotifierPluginSlackWebHook : cannot execute because params is empty")
		return
	}

	webHookURL := ""
	channel := ""

	for key, value := range This.Params {
		if key == "url" {
			if data, ok := value.(string); ok {
				webHookURL = data
			}
		} else if key == "channel" {
			if data, ok := value.(string); ok {
				channel = data
			}
		}
	}

	messageBody := SlackCreateMessage(healthcheck, healthcheckNotifier)
	messageAttachement := SlackCreateAttachement(healthcheck, healthcheckNotifier)

	log.Printf("NotifierPluginSlackWebHook : message to send (to: %v)", messageBody)
	err := This.SendSlack(webHookURL, channel, messageBody, messageAttachement)

	if err != nil {
		log.Printf("NotifierPluginSlackWebHook : executed with error: %v", err)
	} else {
		log.Print("NotifierPluginSlackWebHook : slack sent")
	}
}

func (This *NotifierPluginSlackWebHook) GetName() string {
	return NOTIFIER_PLUGIN_SLACK_WEBHOOK_NAME
}

func (This *NotifierPluginSlackWebHook) GetId() string {
	return This.ID
}

func (This *NotifierPluginSlackWebHook) SendSlack(webHookURL string, channel string, messageBody string, messageAttachement *slack.Attachment) error {
	attachments := []*slack.Attachment{}

	if messageAttachement != nil {
		messageAttachement.Fallback = messageBody
		attachments = append(attachments, messageAttachement)
		messageBody = ""
	}

	hook := slack.NewWebHook(webHookURL)
	err := hook.PostMessage(&slack.WebHookPostPayload{
		Text:        messageBody,
		Channel:     channel,
		Attachments: attachments,
	})

	return err
}
