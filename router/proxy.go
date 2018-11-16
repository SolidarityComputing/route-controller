package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kfcoding-container-api/configs"
	"github.com/kfcoding-container-api/service"
)

// SetupWorkspaceRouter, 设置workspace相关的路由
func SetupWorkspaceRouter(service service.WorkspaceService, r *gin.Engine) {

	var w *gin.RouterGroup

	if configs.AuthPassword == "" && configs.AuthAccount == "" {
		w = r.Group("/proxy")
	} else {
		w = r.Group("/proxy", gin.BasicAuth(gin.Accounts{
			configs.AuthAccount: configs.AuthPassword,
		}))
	}

	w.POST("/add", service.CreateRoutingApi)
	w.POST("/delete", service.CreateRoutingApi)
}
