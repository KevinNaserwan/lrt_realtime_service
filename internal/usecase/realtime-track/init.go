package realtimetrack

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/request"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/response"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/util/env"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/util/redis"
)

type realtimeTrack struct {
}

func NewRealtimeTrack(config *env.Config) Usecase {
	return &realtimeTrack{}
}

// GetRealtimeTrack is a function to get realtime track
func (rt *realtimeTrack) GetRealtimeTrack(ctx context.Context, req request.GeocodingRequest, GoogleAPIKey string) (res response.GeocodingResponse, err error) {
	redisClient := redis.GetRedisClient()
	cacheKey := fmt.Sprintf("geocode:%f:%f", req.Latitude, req.Longitude)

	// Check if data exists in Redis
	cachedData, err := redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Data found in cache, unmarshal and return
		err = json.Unmarshal([]byte(cachedData), &res)
		if err == nil {
			return res, nil
		}
	}

	// Data not found in cache or unmarshaling failed, fetch from API
	res, err = rt.fetchFromAPI(ctx, req, GoogleAPIKey)
	if err != nil {
		return res, err
	}

	// Cache the new data
	jsonData, err := json.Marshal(res)
	if err == nil {
		redisClient.Set(ctx, cacheKey, jsonData, time.Minute) // Cache for 1 minute
	}

	return res, nil
}

func (rt *realtimeTrack) fetchFromAPI(ctx context.Context, req request.GeocodingRequest, GoogleAPIKey string) (response.GeocodingResponse, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?latlng=%f,%f&key=%s",
		req.Latitude, req.Longitude, GoogleAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return response.GeocodingResponse{}, err
	}
	defer resp.Body.Close()

	var apiResponse struct {
		Results []response.GeocodingResponse `json:"results"`
		Status  string                       `json:"status"`
	}

	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return response.GeocodingResponse{}, err
	}

	if len(apiResponse.Results) == 0 {
		return response.GeocodingResponse{}, fmt.Errorf("no results found")
	}

	return apiResponse.Results[0], nil
}

func (rt *realtimeTrack) StartPeriodicUpdate(ctx context.Context, req request.GeocodingRequest, GoogleAPIKey string) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_, err := rt.GetRealtimeTrack(ctx, req, GoogleAPIKey)
			if err != nil {
				fmt.Printf("Error updating real-time data: %v\n", err)
			}
		case <-ctx.Done():
			return
		}
	}
}
