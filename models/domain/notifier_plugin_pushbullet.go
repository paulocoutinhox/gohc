package domain

import (
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
	"log"
)

const (
	NOTIFIER_PLUGIN_PUSHBULLET_NAME = "pushbullet"
)

type NotifierPluginPushBullet struct {
	Params map[string]interface{}
	ID     string
}

func (This *NotifierPluginPushBullet) Notify(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) {
	log.Println("NotifierPluginPushBullet : Notify")

	if This.Params == nil {
		log.Println("NotifierPluginPushBullet : cannot execute because params is empty")
		return
	}

	pbAccessToken := ""
	pbDeviceIden := ""
	pbEmail := ""
	pbChannelTag := ""
	pbClientIden := ""

	for key, value := range This.Params {
		if key == "accessToken" {
			if data, ok := value.(string); ok {
				pbAccessToken = data
			}
		} else if key == "deviceIden" {
			if data, ok := value.(string); ok {
				pbDeviceIden = data
			}
		} else if key == "email" {
			if data, ok := value.(string); ok {
				pbEmail = data
			}
		} else if key == "channelTag" {
			if data, ok := value.(string); ok {
				pbChannelTag = data
			}
		} else if key == "clientIden" {
			if data, ok := value.(string); ok {
				pbClientIden = data
			}
		}
	}

	pushTitle := PushCreateTitle(healthcheck, healthcheckNotifier)
	pushBody := PushCreateMessage(healthcheck, healthcheckNotifier)

	log.Printf("NotifierPluginPushBullet : push to send (body: %v)", pushBody)
	err := This.SendPush(pbAccessToken, pbDeviceIden, pbEmail, pbChannelTag, pbClientIden, pushTitle, pushBody)

	if err != nil {
		log.Printf("NotifierPluginPushBullet : executed with error: %v", err)
	} else {
		log.Print("NotifierPluginPushBullet : push sent")
	}
}

func (This *NotifierPluginPushBullet) GetName() string {
	return NOTIFIER_PLUGIN_PUSHBULLET_NAME
}

func (This *NotifierPluginPushBullet) GetId() string {
	return This.ID
}

func (This *NotifierPluginPushBullet) SendPush(pbAccessToken string, pbDeviceIden string, pbEmail string, pbChannelTag string, pbClientIden string, pushTitle string, pushBody string) error {
	pb := pushbullet.New(pbAccessToken)

	n := requests.NewNote()
	n.Title = pushTitle
	n.Body = pushBody
	n.DeviceIden = pbDeviceIden
	n.Email = pbEmail
	n.ChannelTag = pbChannelTag
	n.ClientIden = pbClientIden

	_, err := pb.PostPushesNote(n)

	return err
}
