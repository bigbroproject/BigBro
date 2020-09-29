package main

import (
	"fmt"
	"github.com/bigbroproject/bigbro/webserver"
	"github.com/bigbroproject/bigbro/webserver/responsehandler"
	"github.com/bigbroproject/bigbrocore/core"
	"github.com/bigbroproject/bigbrocore/responsehandlers"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"log"
	"strconv"
	"time"
)

func main() {

	ws := webserver.NewWebServer("config/serverconfig.yml")
	ws.Start()

	host.Info()

	c, _ := cpu.Info()
	h, _ := host.Info()
	p, _ := cpu.Percent(time.Millisecond*100, false)
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	log.Println("========================")
	log.Println(h.Platform, h.PlatformVersion, h.KernelArch, h.Uptime)
	log.Println(c[0].ModelName, c[0].Mhz, strconv.Itoa(len(c))+" cores", fmt.Sprintf("Usage: %f%%\n", p[0]))
	fmt.Printf("Total: %v, Free:%v, UsedPercent: %f%%\n", v.Total, v.Free, v.UsedPercent)
	//log.Println(float64(memoryInfo.Total - memoryInfo.Available)/(math.Pow(1024, 2)), float64(memoryInfo.Available) / math.Pow(1024, 2), float64(memoryInfo.Total) / (math.Pow(1024, 2)))
	log.Println("========================")

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
