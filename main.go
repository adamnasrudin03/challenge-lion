package main

import (
	"fmt"
	"net/http"

	"adamnasrudin03/challenge-lion/app"
	"adamnasrudin03/challenge-lion/app/configs"
	"adamnasrudin03/challenge-lion/app/controller"
	routers "adamnasrudin03/challenge-lion/app/router"
	"adamnasrudin03/challenge-lion/pkg/database"
	"adamnasrudin03/challenge-lion/pkg/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	dbOne *gorm.DB = database.SetupDBonnectionOne()
	dbTwo *gorm.DB = database.SetupDBonnectionTwo()

	repo     = app.WiringRepository(dbOne, dbTwo)
	services = app.WiringService(repo)

	spController controller.SourceProductController      = controller.NewSourceProductController(services)
	dpController controller.DestinationProductController = controller.NewDestinationProductController(services)
)

func main() {
	defer database.CloseDBConnectionOne(dbOne)
	defer database.CloseDBConnectionTwo(dbTwo)
	config := configs.GetInstance()

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.APIResponse("Welcome my app", http.StatusOK, nil))
	})

	// Route here
	routers.SourceProductRouter(router, spController)
	routers.DestinationProductRouter(router, dpController)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helpers.APIResponse("Page not found", http.StatusNotFound, nil))
	})

	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
