package realtimetrack

import (
	"context"

	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/request"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/response"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/util/env"
)

type realtimeTrack struct {
	config *env.Config
}

func NewRealtimeTrack(config *env.Config) usecase {
	return &realtimeTrack{
		config: config,
	}
}

// GetRealtimeTrack is a function to get realtime track
func (rt *realtimeTrack) GetRealtimeTrack(ctx context.Context, req request.GeocodingRequest, Google_API_Key string) (res response.GeocodingResponse, err error) {

	return res, nil
}
