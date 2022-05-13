package product

import (
	"anaconda/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func (h Handler) handleAddProduct(c *gin.Context) {
	var params ProductRequest
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err := h.service.SubmitProduct(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.Response("success"))
}

func (h Handler) handleGetProduct(c *gin.Context) {
	res, err := h.service.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.Response(res))
}
