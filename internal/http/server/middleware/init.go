package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	httpCommon "github.com/kevinnaserwan/lrt-realtime-service/internal/http/server"
	errorCommon "github.com/kevinnaserwan/lrt-realtime-service/internal/util/error"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors[0]
			// if err can be casted to ClientError, then it is a client error
			if clientError, ok := err.Err.(*errorCommon.ClientError); ok {
				c.JSON(clientError.Code, httpCommon.Error{
					Message: clientError.Message,
				})
			} else if err.IsType(gin.ErrorTypeBind) {
				c.JSON(400, httpCommon.Error{
					Message: err.Err.Error(),
				})
			} else if err.IsType(gin.ErrorTypePrivate) {
				fmt.Println(err.Err.Error())
				c.JSON(500, httpCommon.Error{
					Message: "Internal server error",
				})
			} else {
				fmt.Println(err.Err.Error())
				c.JSON(500, httpCommon.Error{
					Message: "Internal server error",
				})
			}
		}
	}
}
