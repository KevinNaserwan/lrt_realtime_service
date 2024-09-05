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

func NewRealtimeTrack(router *gin.Engine, realtimeTrackUsecase realtimeTrackUsecase.Usecase) {
	controller := &realtimeTrackController{
		RealtimeTrackUsecase: realtimeTrackUsecase,
	}

	router.GET("/realtime-track", controller.GetRealtimeTrack)
}
