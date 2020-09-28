package main

import (
	"github.com/bigbroproject/bigbro/models/data"
	"github.com/bigbroproject/bigbro/webserver"
	"github.com/bigbroproject/bigbro/webserver/responsehandler"
	"github.com/bigbroproject/bigbrocore/core"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
)

func main() {

	servicesList := make([]data.ServiceData, 0)
	servicesListP := &servicesList
	ws := webserver.NewWebServer("config/serverconfig.yml", &servicesListP)
	ws.Start()

	regProtocolInterfaces, regResponseHandlerInterfaces := core.Initialize("config/config.yml")

	// Register custom protocols
	//protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register Response Handlers
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "webServerHandler", responsehandler.WebServerRespHandler{ServicesListP: &servicesListP})
	//responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "console", responsehandlers.ConsoleHandler{})

	// Start monitoring
	core.Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
