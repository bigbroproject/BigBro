package webserver

import (
	"bigbro/models/config"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type WebServer struct {
	Address string
	Port int
	SSL bool
	Router *gin.Engine
}

func NewWebServer(serverConfPath string) *WebServer{

	sConf,err := config.ServerConfigFromFile(serverConfPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	router := gin.Default()

	ws := WebServer{
		Address: sConf.Address,
		Port:    sConf.Port,
		SSL:     sConf.SSL,
		Router:  router,
	}

	// Register REST
	router.Static("/", "./webserver/www/")
	return &ws
}

func (ws *WebServer) Start(){
	go func() {
		err := ws.Router.Run(ws.Address+":"+strconv.Itoa(ws.Port))
		log.Fatal(err.Error())
	}()
}
