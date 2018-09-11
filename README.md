# GoHC

GoHC was made to be a simple and lightweight healthcheck system made with Go (Golang).

Some project features:
- Lightweight (< 5 MB RAM memory)
- Healthcheck list is a simples JSON file - yes, you dont need one database!
- Healthcheck can have 3 types (ping, range and manual)
  - Type "ping": will automatic change status by ping time
  - Type "range": will automatic change status by range values (can be float)
  - Type "manual": will change status using your sent status (work as a trigger)
- Have a warm time configuration, to only start run healthchecks after it (time is in milliseconds and is optional)
- Have a active timeout check for manual and range type (time is in milliseconds and is optional)
- Notification system based on plugins. Today GoHC implements its plugins:
  - CLI
  - Http Get
  - SendGrid
  - PushBullet
  - Slack
- You can reload your healthchecks and notifiers file from web and API - dont need restart the GoHC
- The web interface is nice - made with Bulma and Vue.js
- It is open-source, you can collaborate reporting bugs and upgrading it
- You can DONATE!

You can see healthcheck list sample file with notifications inside: **extras/sample/healthchecks.json**

# Configuration

GoHC configuration is a simple JSON file called "config.json" or what name you want.

Example of configuration file:

```
{
    "server": {
		"host": "0.0.0.0:8080",
		"warmTime": 60000
	},
	"healthchecks": [
	
	],
	"notifiers": [
	
	]
}
```

# Sample files

I have created a sample healthcheck file and a sample config file. Check it on **extras/sample** directory.

# Starting

1. Execute: go get -u github.com/prsolucoes/gohc
2. Execute: cd $GOPATH/src/github.com/prsolucoes/gohc
3. Execute: make deps  
4. Execute: make install  
5. Create config file (config.json) based on some above example  
6. Execute: gohc -f config.json
7. Open in your browser: http://localhost:8080  

# API

1. update a ping = http://localhost:8080/api/update/ping/[TOKEN]
2. update a range = http://localhost:8080/api/update/range/[TOKEN]/[VALUE-FLOAT-OR-INT]
3. update manual = http://localhost:8080/api/update/manual/[TOKEN]/[SUCCESS-WARNING-OR-ERROR]
4. reload healthchecks and notifiers file = http://localhost:8080/api/system/reload

# Command line interface

You can use some make commands to control GoHC service, like start, stop and update from git repository.

1. make stop   = it will kill current GoHC process
2. make update = it will update code from git and install on $GOPATH/bin directory
3. make deps   = download all dependencies
4. make format = format all files (use it before make a pull-request)

So if you want start your server for example, you only need call "make start".

# Alternative method to Build and Start project

1. go build
2. ./gohc -f config.json

# Sugestion

Today, only some functions are implemented. If you need one, you can make a pull-request or send a message in Github Issue.

# Screenshots

**# HOME**

![SS1](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot1.png "Screenshot 1")

**# HEALTHCHECK LIST**

![SS2](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot2.png "Screenshot 2")

**# HEALTHCHECK CHART**

![SS2](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot3.png "Screenshot 3")

**# HEALTHCHECK DASHBOARD - FULL SCREEN**

![SS3](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot4.png "Screenshot 4")

**# HEALTHCHECK DASHBOARD - SMARTPHONE**

<img src="https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot5.png" alt="Screenshot 5" width="240">  

# Supported By Jetbrains IntelliJ IDEA

![Supported By Jetbrains IntelliJ IDEA](https://github.com/prsolucoes/gohc/raw/master/extras/jetbrains/logo.png "Supported By Jetbrains IntelliJ IDEA")

# Author WebSite

> http://www.pcoutinho.com

# License

MIT

# Support with donation
[![Support with donation](http://donation.pcoutinho.com/images/donate-button.png)](http://donation.pcoutinho.com/)
