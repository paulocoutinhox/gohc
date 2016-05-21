# Support with donation
[![Support with donation](http://donation.pcoutinho.com/images/donate-button.png)](http://donation.pcoutinho.com/)

# GoHC

GoHC was made to be a simple and light healthcheck system made with Go (Golang).

Some project features:
- Lightweight (< 5 MB RAM memory)
- Healthcheck list is a simples JSON file - yes, you dont need one database!
- Healthcheck can have 3 types (ping, range and manual)
  - Type "ping": will automatic change status by ping time
  - Type "range": will automatic change status by range values (can be float)
  - Type "status": will change status using your sent status (work as a trigger)
- Have a warm time configuration, to only start run healthchecks after it
- Notification system based on plugins. Today we have CLI, Http Get, SendGrid and PushBullet plugins
- The web interface is nice - made with bootstrap and AJAX
- It is open-source, you can collaborate reporting bugs and upgrading it
- You can DONATE!

# Configuration

GoHC configuration is a simple INI file called "config.ini".

Example of:

```
[server]
host = :8080
workspaceDir = YOUR-WORKSPACE-DIRECTORY
resourcesDir = YOUR-GOPATH-DIRECTORY + /src/github.com/prsolucoes/gohc
warmTime = 10
```

# Sample files

I have created a sample healthcheck file and a sample config file. Check it on **extras/sample** directory.

# Starting

1. Execute: go get -u github.com/prsolucoes/gohc
2. Execute: cd $GOPATH/src/github.com/prsolucoes/gohc
3. Execute: make deps  
4. Execute: make install  
5. Create config file (config.ini) based on some above example  
6. Execute: gohc -f=config.ini
7. Open in your browser: http://localhost:8080  

** dont use character / on any configuration path **

# API

**Check file: controllers/api.go**  
Today we dont have a API doc - but is simple looking code [TODO]  

# Command line interface

You can use some make commands to control GoHC service, like start, stop and update from git repository.

1. make stop   = it will kill current GoHC process
2. make update = it will update code from git and install on $GOPATH/bin directory
3. make deps   = download all dependencies
4. make format = format all files (use it before make a pull-request)

So if you want start your server for example, you only need call "make start".

# Alternative method to Build and Start project

1. go build
2. ./gohc -f=config.ini

# Sugestion

Today, only some functions are implemented. If you need one, you can make a pull-request or send a message in Github Issue.

# Supported By Jetbrains IntelliJ IDEA

![Supported By Jetbrains IntelliJ IDEA](https://github.com/prsolucoes/gohc/raw/master/extras/jetbrains/logo.png "Supported By Jetbrains IntelliJ IDEA")

# Author WebSite

> http://www.pcoutinho.com

# License

MIT

# Screenshots

**# HOME**

![SS1](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot1.png "Screenshot 1")

**# HEALTHCHECK LIST**

![SS2](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot2.png "Screenshot 2")

**# HEALTHCHECK DASHBOARD - FULL SCREEN**

![SS3](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot3.png "Screenshot 3")

**# HEALTHCHECK DASHBOARD - TABLET**

![SS4](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot4.png "Screenshot 4")

**# HEALTHCHECK DASHBOARD - SMARTPHONE**

![SS5](https://github.com/prsolucoes/gohc/raw/master/extras/screenshots/screenshot5.png "Screenshot 5")


