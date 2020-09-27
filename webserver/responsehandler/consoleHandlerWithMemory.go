package responsehandlers

import (
	"fmt"
	"github.com/bigbroproject/bigbrocore/models/response"
	"github.com/fatih/color"
	"log"
	"strconv"
	"time"
)

type WebServerRespHandler struct {
	ServiceProtocol map[string]response.ResponseType
}

func (handler WebServerRespHandler) Handle(channel *chan response.Response) {
	handler.ServiceProtocol = make(map[string]response.ResponseType)
	for {
		resp := <-*channel
		printIfChange(resp, &handler)
	}
}

func printIfChange(resp response.Response, c *WebServerRespHandler) {
	respType := c.ServiceProtocol[resp.ServiceName+resp.Protocol.Type]
	if respType != resp.ResponseType {
		c.ServiceProtocol[resp.ServiceName+resp.Protocol.Type] = resp.ResponseType
		now := time.Now()
		port := strconv.Itoa(resp.Protocol.Port)
		message := ""
		if port == "0" {
			port = "No port"
		}
		if resp.ResponseType == response.Error {
			red := color.New(color.FgRed).SprintFunc()
			message = fmt.Sprintf("[%s] [%s] [%s] [%s - %s - %s] An error as accured: %s", red("ERRO"), now.Format(time.RFC3339), resp.ServiceName, resp.Protocol.Type, resp.Protocol.Server, port, resp.Error.Error())
		} else {
			green := color.New(color.FgHiGreen).SprintFunc()
			message = fmt.Sprintf("[ %s ] [%s] [%s] [%s - %s - %s] Service seems OK.", green("OK"), now.Format(time.RFC3339), resp.ServiceName, resp.Protocol.Type, resp.Protocol.Server, port)
		}
		log.Println(message)
	}
}
