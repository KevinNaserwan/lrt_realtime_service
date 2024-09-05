package request

// GeocodingRequest is a struct to store geocoding request
type GeocodingRequest struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	ResultType string  `json:"result_type"`
	Region     string  `json:"region"`
}
