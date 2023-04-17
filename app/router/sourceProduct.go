package router

import (
	"adamnasrudin03/challenge-lion/app/controller"

	"github.com/gin-gonic/gin"
)

func SourceProductRouter(e *gin.Engine, h controller.SourceProductController) {
	spRoutes := e.Group("/source-products")
	{

		spRoutes.GET("/", h.GetAll)
	}
}
