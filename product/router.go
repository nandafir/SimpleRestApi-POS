package product

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, service *Service) {
	handler := &Handler{
		service: service,
	}
	v1 := router.Group("/v1/anaconda")

	product := v1.Group("/product")
	{
		product.POST("", handler.handleAddProduct)
		product.GET("", handler.handleGetProduct)
	}
}
