PROJECT=gohc
LOG_FILE=/var/log/${PROJECT}.log
GOFMT=gofmt -w
GODEPS=go get -u

.DEFAULT_GOAL := help

# general
help:
	@echo "Type: make [rule]. Available options are:"
	@echo ""
	@echo "- help"
	@echo "- build"
	@echo "- install"
	@echo "- format"
	@echo "- deps"
	@echo "- stop"
	@echo "- start"
	@echo "- update"
	@echo "- build-all"
	@echo ""

build:
	go build -o ${PROJECT}

install:
	go install

format:
	${GOFMT} main.go
	${GOFMT} app/server.go
	${GOFMT} app/binaryFS.go
	${GOFMT} controllers/api.go
	${GOFMT} models/domain/configuration_file.go
	${GOFMT} models/domain/healthcheck.go
	${GOFMT} models/domain/healthcheck_notifier.go
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
	${GOFMT} models/warm/warm.go
	${GOFMT} processor/processor.go
	${GOFMT} template/template.go

deps:
	${GODEPS} github.com/prsolucoes/gowebresponse
	${GODEPS} github.com/gin-gonic/gin
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

	mkdir -p build/linux32
	env GOOS=linux GOARCH=386 go build -o build/linux32/gohc -v github.com/prsolucoes/gohc

	mkdir -p build/linux64
	env GOOS=linux GOARCH=amd64 go build -o build/linux64/gohc -v github.com/prsolucoes/gohc

	mkdir -p build/darwin64
	env GOOS=darwin GOARCH=amd64 go build -o build/darwin64/gohc -v github.com/prsolucoes/gohc

	mkdir -p build/windows32
	env GOOS=windows GOARCH=386 go build -o build/windows32/gohc.exe -v github.com/prsolucoes/gohc

	mkdir -p build/windows64
	env GOOS=windows GOARCH=amd64 go build -o build/windows64/gohc.exe -v github.com/prsolucoes/gohc