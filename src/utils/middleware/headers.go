package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseAllowedOrigins(raw string) []string {
	parts := strings.Split(raw, ",")
	origins := make([]string, 0, len(parts))

	for _, part := range parts {
		origin := strings.TrimSpace(part)
		if origin == "" {
			continue
		}

		origins = append(origins, origin)
	}

	if len(origins) == 0 {
		return []string{"*"}
	}

	return origins
}

func HeadersMiddleware(appName, appVersion string, allowedOrigins []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := generateRequestID()
		origin := ctx.GetHeader("Origin")

		ctx.Header("X-Request-ID", requestID)
		ctx.Header("X-App-Name", appName)
		ctx.Header("X-API-Version", appVersion)
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("X-Frame-Options", "DENY")
		ctx.Header("Referrer-Policy", "no-referrer")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Origin,Content-Type,Accept,Authorization,X-Requested-With")

		if origin == "" {
			ctx.Next()
			return
		}

		if allowedOrigin := matchAllowedOrigin(origin, allowedOrigins); allowedOrigin != "" {
			ctx.Header("Access-Control-Allow-Origin", allowedOrigin)
			ctx.Header("Vary", "Origin")
		}

		if ctx.Request.Method == "OPTIONS" {
			ctx.Status(204)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func matchAllowedOrigin(origin string, allowedOrigins []string) string {
	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == "*" || allowedOrigin == origin {
			return allowedOrigin
		}
	}

	return ""
}

func generateRequestID() string {
	buffer := make([]byte, 12)
	if _, err := rand.Read(buffer); err != nil {
		return "request-id-unavailable"
	}

	return hex.EncodeToString(buffer)
}
