package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sc-route-api/configs"
	"github.com/sc-route-api/router"
	"github.com/sc-route-api/service"
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
