package main

import (
	"github.com/bigbroproject/bigbro/webserver"
	"github.com/bigbroproject/bigbro/webserver/responsehandler"
	"github.com/bigbroproject/bigbrocore/core"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
)

func main() {

	ws := webserver.NewWebServer("config/serverconfig.yml")
	ws.Start()

	regProtocolInterfaces, regResponseHandlerInterfaces := core.Initialize("config/config.yml")

	// Register custom protocols
	//protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register Response Handlers
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "webServerHandler", responsehandler.WebServerRespHandler{OutputChannel: ws.InputChannel})

	//responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "console", responsehandlers.ConsoleHandler{})
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "consoleMemory", responsehandlers.ConsoleHandlerWithMemory{})

	// Start monitoring
	core.Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
