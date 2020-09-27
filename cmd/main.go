package main

import (
	"bigbro/webserver"
	"github.com/bigbroproject/bigbrocore/core"
	"github.com/bigbroproject/bigbrocore/protocols"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
)

func main() {

	ws := webserver.NewWebServer("config/serverconfig.yml")
	ws.Start()

	regProtocolInterfaces, regResponseHandlerInterfaces := core.Initialize("config/config.yml")

	// Register custom protocol
	protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register custom Response Handler
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "consoleWithMemory", responsehandlers.ConsoleHandlerWithMemory{})
	//responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "console", responsehandlers.ConsoleHandler{})

	// Start monitoring
	core.Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
