package realtimetrack

import (
	"context"

	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/request"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/response"
)

type Usecase interface {
	GetRealtimeTrack(ctx context.Context, req request.GeocodingRequest, Google_API_Key string) (res response.GeocodingResponse, err error)
	fetchFromAPI(ctx context.Context, req request.GeocodingRequest, GoogleAPIKey string) (response.GeocodingResponse, error)
	StartPeriodicUpdate(ctx context.Context, req request.GeocodingRequest, GoogleAPIKey string)
}
