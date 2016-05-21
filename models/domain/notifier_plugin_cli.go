package domain

import (
	"log"
	"os/exec"
)

const (
	NOTIFIER_PLUGIN_CLI_NAME = "cli"
)

type NotifierPluginCLI struct {
	Params map[string]interface{}
	ID     string
}

func (This *NotifierPluginCLI) Notify(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) {
	log.Println("NotifierPluginCLI : Notify")

	if This.Params == nil {
		log.Println("NotifierCLI : cannot execute because params is empty")
		return
	}

	command := ""
	args := []string{}
	workingDir := ""

	for key, value := range This.Params {
		if key == "workingDir" {
			if data, ok := value.(string); ok {
				workingDir = data
			}
		} else if key == "command" {
			if data, ok := value.(string); ok {
				command = data
			}
		} else if key == "args" {
			if data, ok := value.([]interface{}); ok {
				for _, arg := range data {
					if argData, ok := arg.(string); ok {
						args = append(args, argData)
					}
				}
			}
		}
	}

	log.Printf("NotifierCLI : command to be execute (command: %v, working dir: %v, args: %v)", command, workingDir, args)

	cmd := exec.Command(command, args...)
	cmd.Dir = workingDir
	out, err := cmd.Output()

	if err != nil {
		log.Printf("NotifierCLI : executed with error: %v", err)
	} else {
		log.Printf("NotifierCLI : executed with output: \n %v", string(out))
	}
}

func (This *NotifierPluginCLI) GetName() string {
	return NOTIFIER_PLUGIN_CLI_NAME
}

func (This *NotifierPluginCLI) GetId() string {
	return This.ID
}
