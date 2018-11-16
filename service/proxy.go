package service

import (
	"github.com/gin-gonic/gin"
	"github.com/kfcoding-container-api/model"
	"net/http"
)

type WorkspaceService interface {
	CreateRoutingApi(c *gin.Context)
	DeleteRoutingApi(c *gin.Context)
}

type ProxyService struct {
	routingService RoutingService
	etcdService    *EtcdService
}

func NewProxyService(routing RoutingService, etcdService *EtcdService) (*ProxyService) {
	return &ProxyService{
		routingService: routing,
		etcdService:    etcdService,
	}
}

func (service *ProxyService) CreateRoutingApi(c *gin.Context) {

	w := &model.AddRoutingBody{}
	if err := c.ShouldBindJSON(w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, v := range w.RoutingList {
		err := service.routingService.AddRule(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"data": "ok"})
	return
}

func (service *ProxyService) DeleteRoutingApi(c *gin.Context) {
	w := &model.DelRoutingBody{}
	if err := c.ShouldBindJSON(w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, v := range w.RoutingList {
		err := service.routingService.DeleteRule(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"data": "ok"})
	return
}
