PROJECT=gohc
LOG_FILE=/var/log/${PROJECT}.log
GOFMT=gofmt -w
GODEPS=go get

GOFILES=\
	main.go\

build:
	go build -o ${PROJECT}

install:
	go install

format:
	${GOFMT} main.go
	${GOFMT} app/server.go
	${GOFMT} controllers/api.go
	${GOFMT} controllers/dashboard.go
	${GOFMT} controllers/healthcheck.go
	${GOFMT} controllers/home.go
	${GOFMT} models/domain/healthcheck.go
	${GOFMT} models/domain/healthcheck_notifier.go
	${GOFMT} models/domain/healthchecks_file.go
	${GOFMT} models/domain/mail.go
	${GOFMT} models/domain/notifier.go
	${GOFMT} models/domain/notifier_plugin_cli.go
	${GOFMT} models/domain/notifier_plugin_http_get.go
	${GOFMT} models/domain/notifier_plugin_interface.go
	${GOFMT} models/domain/notifier_plugin_manager.go
	${GOFMT} models/domain/notifier_plugin_pushbullet.go
	${GOFMT} models/domain/notifier_plugin_sendgrid.go
	${GOFMT} models/domain/notifier_plugin_slack_webhook.go
	${GOFMT} models/domain/push.go
	${GOFMT} models/domain/slack.go
	${GOFMT} models/util/util.go
	${GOFMT} processor/processor.go

test:

deps:
	${GODEPS} github.com/prsolucoes/gowebresponse
	${GODEPS} github.com/gin-gonic/gin
	${GODEPS} github.com/go-ini/ini
	${GODEPS} github.com/sendgrid/sendgrid-go
	${GODEPS} github.com/mitsuse/pushbullet-go
	${GODEPS} github.com/bluele/slack

stop:
	pkill -f ${PROJECT}

start:
	-make stop
	cd ${GOPATH}/src/github.com/prsolucoes/${PROJECT}
	nohup ${PROJECT} >> ${LOG_FILE} 2>&1 </dev/null &

update:
	git pull origin master
	make install