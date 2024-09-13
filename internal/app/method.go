package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	docs "github.com/kevinnaserwan/lrt-realtime-service/api"
	realtimeTrackController "github.com/kevinnaserwan/lrt-realtime-service/internal/controller/realtime-track"
	http "github.com/kevinnaserwan/lrt-realtime-service/internal/http/server"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/server/middleware"
	RealtimeTrackUsecase "github.com/kevinnaserwan/lrt-realtime-service/internal/usecase/realtime-track"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/util/env"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *app) StartServer() {
	server := http.NewHTTPServer(a.config.Environment)

	server.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"message": "Welcome to lrt sumsel realtime track service",
		})
	})

	if a.config.Environment != "release" {
		log.Println("Initializing Swagger...")
		docs.SwaggerInfo.Title = "LRT SUMSEL Realtime Track API"
		docs.SwaggerInfo.Description = "API Documentation for LRT SUMSEL Realtime Track Service"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", a.config.Port)
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}

		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Println("Swagger initialized successfully")
	}

	initController(server, a.config)

	log.Printf("Server is running on port %d", a.config.Port)
	err := server.Run(fmt.Sprintf(":%d", a.config.Port))
	if err != nil {
		panic(err)
	}

}

func initController(root *gin.Engine, config *env.Config) {
	// Initialize the real-time track usecase with the config
	realtimeTrackUsecase := RealtimeTrackUsecase.NewRealtimeTrack(config)

	routerGroup := root.Group("/api/v1")
	routerGroup.Use(middleware.ErrorHandler())

	// Initialize the real-time track controller and route, pass config to controller
	realtimeTrackController.NewRealtimeTrackController(routerGroup.Group("/track"), realtimeTrackUsecase, config)
}
