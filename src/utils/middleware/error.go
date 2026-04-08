package middleware

import (
	"go-api/src/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddlewareHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			err := context.Errors.Last().Err

			apiErr, ok := err.(*dtos.APIError)

			if ok {
				context.JSON(apiErr.StatusCode, gin.H{"error": apiErr.Message})
				context.Abort()
				return
			}
			context.JSON(http.StatusInternalServerError, dtos.APIError{StatusCode: http.StatusInternalServerError, Message: err.Error()})
			context.Abort()
		}
	}
}
