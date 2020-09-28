package webserver

import (
	"github.com/bigbroproject/bigbro/models/config"
	"github.com/bigbroproject/bigbro/models/data"
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
	ServicesList **[]data.ServiceData
}

func NewWebServer(serverConfPath string, servicesList **[]data.ServiceData) *WebServer {

	sConf, err := config.ServerConfigFromFile(serverConfPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	router := gin.Default()

	ws := WebServer{
		Address:      sConf.Address,
		Port:         sConf.Port,
		SSL:          sConf.SSL,
		Router:       router,
		ServicesList: servicesList,
	}

	// Register REST
	staticGroup := router.Group("/dashboard")
	staticGroup.Static("/", "./webserver/www/")

	apiGroup := router.Group("/api")
	apiGroup.GET("/services", func(context *gin.Context) { getServicesList(context, &ws) })

	router.GET("/", func(context *gin.Context) { context.Redirect(http.StatusMovedPermanently, "/dashboard") })

	return &ws
}

func (ws *WebServer) Start() {
	go func() {
		err := ws.Router.Run(ws.Address + ":" + strconv.Itoa(ws.Port))
		log.Fatal(err.Error())
	}()
}

func getServicesList(context *gin.Context, ws *WebServer) {

	context.JSON(http.StatusOK, ws.ServicesList)

}
