package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kfcoding-container-api/configs"
	"github.com/kfcoding-container-api/router"
	"github.com/kfcoding-container-api/service"
	"log"
)

func main() {

	// init
	configs.InitEnvs()
	gin.SetMode(gin.ReleaseMode)

	// new services
	etcdService := service.NewEtcdService()
	routingService := service.NewRoutingTraefikService(service.NewEtcdService())
	workspaceService := service.NewProxyService(routingService, etcdService)

	// setup routers
	r := gin.Default()
	router.SetupWorkspaceRouter(workspaceService, r)

	// start server
	log.Print("---> Serve on ", configs.ServeAddress)
	r.Run(configs.ServeAddress)
}
