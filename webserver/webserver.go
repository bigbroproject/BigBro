package webserver

import (
	"fmt"
	"github.com/bigbroproject/bigbro/models/config"
	"github.com/bigbroproject/bigbro/models/data"
	"github.com/bigbroproject/bigbro/system"
	"github.com/bigbroproject/bigbrocore/models/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"strconv"
	"github.com/rakyll/statik/fs"

	_ "github.com/bigbroproject/bigbro/webserver/statik"
)

const (
	ROOM_SERVICES_LISTENERS = "serviceListeners"
)

const (
	EVENT_SERVICE_CHANGE = "serviceChange"
)

type WebServer struct {
	Address      string
	Port         int
	SSL          bool
	Router       *gin.Engine
	InputChannel chan response.Response
	ServiceMap   *map[string]data.ServiceData
	ServerSocket *socketio.Server
}

func NewWebServer(serverConfPath string) *WebServer {

	sConf, err := config.ServerConfigFromFile(serverConfPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	serverSocket := newServerSocket()
	go func() {
		err := serverSocket.Serve()

		if err != nil {
			log.Fatal(err)
		}
	}()

	//defer serverSocket.Close()

	gin.SetMode(gin.ReleaseMode)

	router := gin.New() // gin.Default()

	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = sConf.AllowOrigins
	configCors.AllowCredentials = true
	router.Use(cors.New(configCors))
	//router.Use(GinMiddleware("localhost:8080"))

	//router.Use(cors.Default())
	//router.Use(CORSMiddleware())
	//router.Use(gin.Logger())
	router.Use(gin.Recovery())
	inputChannel := make(chan response.Response)
	serviceMap := make(map[string]data.ServiceData, 0)

	ws := WebServer{
		Address:      sConf.Address,
		Port:         sConf.Port,
		SSL:          sConf.SSL,
		Router:       router,
		ServiceMap:   &serviceMap,
		InputChannel: inputChannel,
		ServerSocket: serverSocket,
	}

	// Register REST
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	staticGroup := router.Group("/dashboard")
	staticGroup.StaticFS("/", statikFS)

	apiGroup := router.Group("/api")
	apiGroup.GET("/services", func(context *gin.Context) { getServicesList(context, &ws) })
	apiGroup.GET("/system", func(context *gin.Context) { getSystemInformation(context) })

	router.GET("/socket.io/*any", gin.WrapH(serverSocket))
	router.POST("/socket.io/*any", gin.WrapH(serverSocket))

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
			sendToWs := false

			for i := 0; i < len(serviceData.Protocols); i++ {
				protocolData := &serviceData.Protocols[i]
				if protocolData.Protocol.Type == resp.Protocol.Type && protocolData.Protocol.Server == resp.Protocol.Server && protocolData.Protocol.Port == resp.Protocol.Port {

					if protocolData.Err != nil && resp.Error != nil {
						if protocolData.Err.Error() != resp.Error.Error() {
							protocolData.Err = resp.Error
							sendToWs = true
						}
					} else {
						if protocolData.Err != resp.Error {
							protocolData.Err = resp.Error
							sendToWs = true
						}
					}
					protocolExists = true
				}

			}

			if !protocolExists {
				serviceData.Protocols = append(serviceData.Protocols, data.ProtocolData{
					Protocol: resp.Protocol,
					Err:      resp.Error,
				})

			}

			if sendToWs {
				ws.ServerSocket.BroadcastToRoom("/", ROOM_SERVICES_LISTENERS, EVENT_SERVICE_CHANGE, serviceData)
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
			auxData := data.ServiceData{
				Name:      resp.ServiceName,
				Protocols: protocolsList,
			}
			(*ws.ServiceMap)[resp.ServiceName] = auxData
		}
	}

}

func newServerSocket() *socketio.Server {

	serverSocket, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	serverSocket.OnConnect("/", func(s socketio.Conn) error {
		s.Join(ROOM_SERVICES_LISTENERS)
		s.SetContext("")
		fmt.Println("connected:", serverSocket.Count())
		return nil
	})

	serverSocket.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	serverSocket.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	serverSocket.OnDisconnect("/", func(s socketio.Conn, reason string) {
		s.Close()
		fmt.Println("closed", reason)
		fmt.Println("connected remain:", serverSocket.Count())
	})

	return serverSocket
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

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}
