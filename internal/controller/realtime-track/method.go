package realtimetrack

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/http/request"
)

// @Summary		Get Realtime Track
// @Description	Get Realtime Track
// @Tags			Realtime Track
// @Accept			json
// @Produce		json
// @Param			request	body		request.GeocodingRequest	true	"Geocoding Request"
// @Success		200		{object}	http.Response{value=response.GeocodingResponse}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/realtime-track [get]
func (c *realtimeTrackController) GetRealtimeTrack(ctx *gin.Context) {
	var req request.GeocodingRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.RealtimeTrackUsecase.GetRealtimeTrack(ctx, req, c.Config.Google_API_Key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
