package master

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, service *Service) {
	handler := &Handler{
		service: service,
	}
	v1 := router.Group("/v1/master")

	//table item
	item := v1.Group("/item")
	{
		item.POST("", handler.handlerAddItem)
		item.GET("", handler.handlerGetItem)
		item.DELETE(":id", handler.handlerSoftDeleteItem)
	}

}
