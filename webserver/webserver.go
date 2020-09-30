package webserver

import (
	"github.com/bigbroproject/bigbro/models/config"
	"github.com/bigbroproject/bigbro/models/data"
	"github.com/bigbroproject/bigbro/system"
	"github.com/bigbroproject/bigbrocore/models/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type WebServer struct {
	Address      string
	Port         int
	SSL          bool
	Router       *gin.Engine
	InputChannel chan response.Response
	ServiceMap   *map[string]data.ServiceData
}

func NewWebServer(serverConfPath string) *WebServer {

	sConf, err := config.ServerConfigFromFile(serverConfPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	inputChannel := make(chan response.Response)
	serviceMap := make(map[string]data.ServiceData, 0)

	ws := WebServer{
		Address:      sConf.Address,
		Port:         sConf.Port,
		SSL:          sConf.SSL,
		Router:       router,
		ServiceMap:   &serviceMap,
		InputChannel: inputChannel,
	}

	// Register REST

	staticGroup := router.Group("/dashboard")
	staticGroup.Static("/", "./webserver/www/")

	apiGroup := router.Group("/api")
	apiGroup.GET("/services", func(context *gin.Context) { getServicesList(context, &ws) })
	apiGroup.GET("/system", func(context *gin.Context) { getSystemInformation(context) })

	router.GET("/", func(context *gin.Context) { context.Redirect(http.StatusMovedPermanently, "/dashboard") })

	return &ws
}

func (ws *WebServer) Start() {
	go func() {
		err := ws.Router.Run(ws.Address + ":" + strconv.Itoa(ws.Port))
		log.Fatal(err.Error())
	}()

	go func() {
		for {
			ws.listenInputChannel()
		}
	}()
}

func (ws *WebServer) listenInputChannel() {

	select {
	case resp := <-ws.InputChannel:
		if serviceData, exists := (*ws.ServiceMap)[resp.ServiceName]; exists {
			//service exists
			protocolExists := false
			for _, protocolData := range serviceData.Protocols {
				if protocolData.Protocol.Type == resp.Protocol.Type && protocolData.Protocol.Server == resp.Protocol.Server && protocolData.Protocol.Port == resp.Protocol.Port {
					protocolData.Err = resp.Error
					protocolExists = true
				}

			}

			if !protocolExists {
				serviceData.Protocols = append(serviceData.Protocols, data.ProtocolData{
					Protocol: resp.Protocol,
					Err:      resp.Error,
				})

			}

			(*ws.ServiceMap)[resp.ServiceName] = serviceData

		} else {
			//service does not exist
			firstProtocol := data.ProtocolData{
				Protocol: resp.Protocol,
				Err:      resp.Error,
			}
			protocolsList := make([]data.ProtocolData, 1)
			protocolsList[0] = firstProtocol
			(*ws.ServiceMap)[resp.ServiceName] = data.ServiceData{
				Name:      resp.ServiceName,
				Err:       resp.Error,
				Protocols: protocolsList,
			}
		}
	}

}

func getServicesList(context *gin.Context, ws *WebServer) {
	context.JSON(http.StatusOK, ws.ServiceMap)
}

func getSystemInformation(context *gin.Context) {
	sysInfo, err := system.GetSystemInfo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	} else {
		context.JSON(http.StatusOK, sysInfo)
	}
}
