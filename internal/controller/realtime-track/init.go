package realtimetrack

import (
	"github.com/gin-gonic/gin"
	realtimeTrackUsecase "github.com/kevinnaserwan/lrt-realtime-service/internal/usecase/realtime-track"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/util/env"
)

type realtimeTrackController struct {
	RealtimeTrackUsecase realtimeTrackUsecase.Usecase
	Config               *env.Config
}

func NewRealtimeTrackController(router *gin.RouterGroup, realtimeTrackUsecase realtimeTrackUsecase.Usecase, config *env.Config) {
	controller := &realtimeTrackController{
		RealtimeTrackUsecase: realtimeTrackUsecase,
		Config:               config, // Pass the config with Google API Key
	}

	router.GET("/realtime-track", controller.GetRealtimeTrack)
}
