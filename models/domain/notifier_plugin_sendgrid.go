package domain

import (
	"github.com/sendgrid/sendgrid-go"
	"log"
)

const (
	NOTIFIER_PLUGIN_SENDGRID_NAME = "sendgrid"
)

type NotifierPluginSendGrid struct {
	Params map[string]interface{}
	ID     string
}

func (This *NotifierPluginSendGrid) Notify(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) {
	log.Println("NotifierPluginSendGrid : Notify")

	if This.Params == nil {
		log.Println("NotifierPluginSendGrid : cannot execute because params is empty")
		return
	}

	mailToList := []string{}
	mailFromEmail := ""
	mailFromName := ""
	sendGridKey := ""

	for key, value := range This.Params {
		if key == "to" {
			if data, ok := value.([]interface{}); ok {
				for _, mailTo := range data {
					if mailToData, ok := mailTo.(string); ok {
						mailToList = append(mailToList, mailToData)
					}
				}
			}
		} else if key == "key" {
			if data, ok := value.(string); ok {
				sendGridKey = data
			}
		} else if key == "fromEmail" {
			if data, ok := value.(string); ok {
				mailFromEmail = data
			}
		} else if key == "fromName" {
			if data, ok := value.(string); ok {
				mailFromName = data
			}
		}
	}

	mailSubject := MailCreateSubject(healthcheck, healthcheckNotifier)
	mailBody := MailCreateBody(healthcheck, healthcheckNotifier)

	log.Printf("NotifierPluginSendGrid : mail to send (to: %v)", mailToList)
	err := This.SendEmail(mailSubject, mailBody, mailFromEmail, mailFromName, mailToList, sendGridKey)

	if err != nil {
		log.Printf("NotifierPluginSendGrid : executed with error: %v", err)
	} else {
		log.Print("NotifierPluginSendGrid : mail sent")
	}
}

func (This *NotifierPluginSendGrid) GetName() string {
	return NOTIFIER_PLUGIN_SENDGRID_NAME
}

func (This *NotifierPluginSendGrid) GetId() string {
	return This.ID
}

func (This *NotifierPluginSendGrid) SendEmail(mailSubject string, mailBody string, mailFromEmail string, mailFromName string, mailToList []string, sendGridKey string) error {
	sg := sendgrid.NewSendGridClientWithApiKey(sendGridKey)

	message := sendgrid.NewMail()

	for _, mailTo := range mailToList {
		message.AddTo(mailTo)
	}

	message.SetFrom(mailFromEmail)
	message.SetFromName(mailFromName)

	message.SetSubject(mailSubject)
	message.SetHTML(mailBody)

	return sg.Send(message)
}
