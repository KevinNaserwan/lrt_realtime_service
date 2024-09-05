package realtimetrack

import (
	"context"

	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/request"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/response"
)

type usecase interface {
	GetRealtimeTrack(ctx context.Context, req request.GeocodingRequest, Google_API_Key string) (res response.GeocodingResponse, err error)
}
