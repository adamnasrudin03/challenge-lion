package controller

import (
	"adamnasrudin03/challenge-lion/app/service"
	"adamnasrudin03/challenge-lion/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type SourceProductController interface {
	GetAll(ctx *gin.Context)
}

type spHandler struct {
	Service *service.Services
}

func NewSourceProductController(srv *service.Services) SourceProductController {
	return &spHandler{
		Service: srv,
	}
}

func (c *spHandler) GetAll(ctx *gin.Context) {
	res, httpStatus, err := c.Service.SourceProduct.GetAll(ctx)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), httpStatus, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", httpStatus, res))
}
