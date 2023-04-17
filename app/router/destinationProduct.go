package router

import (
	"adamnasrudin03/challenge-lion/app/controller"

	"github.com/gin-gonic/gin"
)

func DestinationProductRouter(e *gin.Engine, h controller.DestinationProductController) {
	dpRoutes := e.Group("/destination-products")
	{

		dpRoutes.GET("/", h.GetAll)
	}
}
