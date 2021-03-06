package main

import (
	"github.com/bigbroproject/bigbro/system"
	"github.com/bigbroproject/bigbro/webserver"
	"github.com/bigbroproject/bigbro/webserver/responsehandler"
	"github.com/bigbroproject/bigbrocore/core"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
	"github.com/bigbroproject/bigbrocore/utilities"
	"github.com/fatih/color"
	"log"
	"os"
)

func main() {

	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	errEnv := os.Setenv("GHW_DISABLE_WARNINGS", "1")
	if errEnv != nil{
		log.Printf("[%s] %s", utilities.CreateColorString("Warning",color.FgHiYellow), errEnv)
	}
	system.PrintSystemInfo()
	ws := webserver.NewWebServer("config/serverconfig.yml")
	ws.Start()

	regProtocolInterfaces, regResponseHandlerInterfaces := core.Initialize("config/config.yml")

	// Register custom protocols
	//protocols.RegisterProtocolInterface(&regProtocolInterfaces, "ftp", protocols.FTP{})

	// Register Response Handlers
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "webServerHandler", responsehandler.WebServerRespHandler{OutputChannel: ws.InputChannel})

	//responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "console", responsehandlers.ConsoleHandler{})
	responsehandlers.RegisterResponseHandlerInterface(&regResponseHandlerInterfaces, "consoleMemory", responsehandlers.ConsoleHandlerWithMemory{})

	//go func() {
	//	for  {
	//		fmt.Println("====== GoRoutine "+strconv.Itoa(runtime.NumGoroutine()))
	//		time.Sleep(time.Second * 5)
	//	}
	//}()
	// Start monitoring
	core.Start(regProtocolInterfaces, regResponseHandlerInterfaces)

}
