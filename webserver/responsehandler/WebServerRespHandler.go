package responsehandler

import (
	"fmt"
	"github.com/bigbroproject/bigbrocore/models"
	"github.com/bigbroproject/bigbrocore/models/response"
	"github.com/fatih/color"
	"log"
	"strconv"
	"time"
)

type WebServerRespHandler struct {
	ServiceProtocol map[string]response.ResponseType
	OutputChannel   chan response.Response
}

func (handler WebServerRespHandler) Handle(configuration *models.Config, channel *chan response.Response) {

	handler.loadServices(configuration)
	handler.ServiceProtocol = make(map[string]response.ResponseType)

	for {
		select {
		case resp := <-*channel:
			handler.writeResponse(resp)
		}

	}
}

func (handler WebServerRespHandler) writeResponse(response response.Response) {
	go func() {
		select {
		case handler.OutputChannel <- response:
			return
		}
	}()
}

func (handler WebServerRespHandler) loadServices(configuration *models.Config) {

	for _, service := range configuration.Services {

		for _, protocol := range service.Protocols {
			resp := response.Response{
				ServiceName: service.Name,
				Protocol:    protocol,
			}
			handler.writeResponse(resp)

		}

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
