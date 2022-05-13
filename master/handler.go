package master

import (
	"anaconda/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func (h Handler) handlerAddItem(c *gin.Context) {
	var params ItemRequest

	fmt.Println("param : ", params)

	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err := h.service.SubmitItem(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.Response("success"))
}

func (h Handler) handlerGetItems(c *gin.Context) {
	res, err := h.service.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.Response(res))
}
