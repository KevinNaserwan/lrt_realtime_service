package response

// GeocodingResponse is a struct to store geocoding response
type GeocodingResponse struct {
	AddressComponents []AddressResponse `json:"address_components"`
	FormattedAddress  string            `json:"formatted_address"`
	Geometry          GeometryResponse  `json:"geometry"`
	PlaceID           string            `json:"place_id"`
	Types             []string          `json:"types"`
}

// AddressResponse is a struct to store address response
type AddressResponse struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// GeometryResponse is a struct to store geometry response
type GeometryResponse struct {
	Location     LocationResponse `json:"location"`
	LocationType string           `json:"location_type"`
	Viewport     ViewportResponse `json:"viewport"`
}

// LocationResponse is a struct to store location response
type LocationResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ViewportResponse is a struct to store viewport response
type ViewportResponse struct {
	Northeast LocationResponse `json:"northeast"`
	Southwest LocationResponse `json:"southwest"`
}
