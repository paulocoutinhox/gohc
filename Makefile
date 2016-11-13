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
	${GOFMT} models/warm/warm.go
	${GOFMT} processor/processor.go
	${GOFMT} app/binaryFS.go
	${GOFMT} template/template.go

test:

deps:
	${GODEPS} github.com/prsolucoes/gowebresponse
	${GODEPS} github.com/gin-gonic/gin
	${GODEPS} github.com/go-ini/ini
	${GODEPS} github.com/sendgrid/sendgrid-go
	${GODEPS} github.com/mitsuse/pushbullet-go
	${GODEPS} github.com/bluele/slack
	${GODEPS} github.com/elazarl/go-bindata-assetfs

stop:
	pkill -f ${PROJECT}

start:
	-make stop
	cd ${GOPATH}/src/github.com/prsolucoes/${PROJECT}
	nohup ${PROJECT} >> ${LOG_FILE} 2>&1 </dev/null &

update:
	git pull origin master
	make deps
	make install

build-all:
	rm -rf build
	mkdir -p build/darwin64
	env GOOS=darwin GOARCH=amd64 go build -o build/darwin64/gohc -v github.com/prsolucoes/gohc
	mkdir -p build/windows32
	env GOOS=windows GOARCH=386 go build -o build/windows32/gohc -v github.com/prsolucoes/gohc
	mkdir -p build/windows64
	env GOOS=windows GOARCH=amd64 go build -o build/windows64/gohc -v github.com/prsolucoes/gohc