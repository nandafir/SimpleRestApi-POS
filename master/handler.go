package master

import (
	"anaconda/utils"
	"fmt"
	"net/http"
	"strconv"

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

func (h Handler) handlerGetItem(c *gin.Context) {
	res, err := h.service.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.Response(res))
}

func (h Handler) handlerSoftDeleteItem(c *gin.Context) {

	idItem, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	err = h.service.SoftDeleteItemByID(idItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, utils.Response(SuccessStatus))
}
