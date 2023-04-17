package controller

import (
	"adamnasrudin03/challenge-lion/app/service"
	"adamnasrudin03/challenge-lion/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type DestinationProductController interface {
	GetAll(ctx *gin.Context)
}

type dpHandler struct {
	Service *service.Services
}

func NewDestinationProductController(srv *service.Services) DestinationProductController {
	return &dpHandler{
		Service: srv,
	}
}

func (c *dpHandler) GetAll(ctx *gin.Context) {
	res, httpStatus, err := c.Service.DestinationProduct.GetAll(ctx)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", httpStatus, res))
}
