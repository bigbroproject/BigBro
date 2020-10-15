<img src="readme/bigbro-logo.svg" width="350">

# BigBro 
[![Build Status](https://travis-ci.org/bigbroproject/bigbrocore.svg?branch=master)](https://travis-ci.org/bigbroproject/bigbrocore) [![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://https://github.com/bigbroproject/bigbrocore/master/LICENSE)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://golang.org)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Vue%203-%2341b883)](https://https://v3.vuejs.org/)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bigbroproject/bigbrocore.svg)](https://github.com/bigbroproject/bigbrocore)
[![GitHub Release](https://img.shields.io/github/release/bigbroproject/bigbro.svg?style=flat)]()  

An extensible monitoring tool for user defined services and protocols, all wrote in GoLang and UI with VueJS 3, made with [BigBro core](http://github.com/bigbroproject/bigbrocore).

### Project status
* This project is in **BETA stage**
* Go version : **1.15**
* Vue version : **3.0**
* BigBro core version: [![GitHub Release](https://img.shields.io/github/release/bigbroproject/bigbrocore.svg?style=flat)]()  

## Introduction

BigBro is a lightweight monitoring software tool that offers a lightweight User Interface to monitor the services and protocols defined by users.

<img src="readme/dashboard.png" >

The software also can be used only via CLI interface

<img src="readme/cli.png" >

The entire software is based on one configuration file where it is possible define a list of services that would be monitor.

## Compatibility

### Operative Systems
BigBro was compiled and test on Debian and RHEL Linux distribution families.

### Hardware
BigBro was develop to run in any kind of computer or embedded system and was tested on:

- General Personal Computer (x86_64)
- Raspberry 3 B+
- Orange Pi Zero

## Defaults Protocols implemented
The default protocols are based on default protocols from BigBro Core, nevertheless new or custom protocols can be implemeted in BigBro as in BigBro Core.
- http
- https (with ssl)
- icmp
- icmp6

## Configuration file
In the configuration file will must be defined all services and protocol to be monitor. The format type of the file is YAML as follows:

```yaml
services:
  - name: Facebook (ssl)
    protocols:
      - type: https
        port: 443
        server: facebook.com
        interval : 1000
  - name: Google
    protocols:
      - type: http
        server: google.com
        interval : 1000
  - name: Google DNS
    protocols:
      - type: icmp
        server: 8.8.8.8
        interval : 1000
  - name: Google DNS6
    protocols:
      - type: icmp6
        server: 2a00:1450:400a:804::2004
        interval : 5000
      - type: ftp
        server: 1.1.1.1
        interval : 1000
        customs :
          user : uss
          password : passwd
  ...
```
* `name` : Name of the service
* `protocols`: List of protocols that would be monitored
    * `type`: name of registered protocol
    * `server`: IP or dns of service
    * `interval`: interval time between check (in milliseconds)
    * `port`: (optional) port of service
    * `customs`: (optional) list of custom filed for custom implementation of protocols


## Usage Example
### From Binary File
Basically BigBro can be downloaded from releases section with a pre-compiled binary file.

[Releases Page](https://github.com/bigbroproject/bigbro/releases) | Last release:

[![GitHub Release](https://img.shields.io/github/release/bigbroproject/bigbro.svg?style=flat)]()  

### From Source Code

#### Requirements
- NodeJS >= v12.10.1 (it is used to compile the user interface)
- Go Version >= v1.15 (developed with this version)
- Git installed
- Linux Debian or RHEL distribution

Firstly, before the usage of this library you must clone this repository:
```bash
git clone github.com/bigbroproject/bigbro
```

Secondly, the installation of libraries and dependencies is required:
```bash
cd bigbro && make install-dep
```
After the installation of dependencies, you can proceed to build BigBro:

```bash
make build-x64
make build-x86
make build-arm
make build-arm-7
```
If there are any errors you can find the built solution into the new dir bin or you can install automatically the solution in your system with install command (BigBro will be installed in your home):

```bash
make install
```


#### Customisation of the code
Before the build phse you can edit the main code, registering new protocols or handlers:

```go
package main

import (
	"github.com/bigbroproject/bigbrocore/core"
	"github.com/bigbroproject/bigbrocore/protocols"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
)

func main() {


	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	system.PrintSystemInfo()
	ws := webserver.NewWebServer("config/serverconfig.yml")
	ws.Start()

	regProtocolInterfaces, regResponseHandlerInterfaces := core.Initialize("config/config.yml")

	// Register custom protocols
	//protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register Response Handlers
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "webServerHandler", responsehandler.WebServerRespHandler{OutputChannel: ws.InputChannel})

    // Register custom handlers
	//responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "console", responsehandlers.ConsoleHandler{})
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "consoleMemory", responsehandlers.ConsoleHandlerWithMemory{})

	// Start monitoring
	core.Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
```
In this example we `Initialize` the entire module with a given configuration file path (you can mange this as a first input in your command line, for instance). 
Secondly, you must define and register a ResponseHandler to manage the Responses from service checks (otherwise you cannot log or see anything) and then register a custom protocol, if needed.

Finally, you can over your project main with the `Start` of the module. 

## Struct of the Core and implementation of custom protocols and handlers

For more information about the implementation of new protocols and handlers read the documentation of [BigBro core](http://github.com/bigbroproject/bigbrocore).


---
Made with ❤️ by [Asdrubbalo](http://github.com/danielemlu), [filirnd](http://github.com/filirnd) and [fedyfausto](http://github.com/fedyfausto)


