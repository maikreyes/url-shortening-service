package cors

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {

	allowedOrigins := parseAllowedOrigins(os.Getenv("CORS_ALLOWED_ORIGINS"))

	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")

		if origin != "" {
			if isOriginAllowed(origin, allowedOrigins) {

				ctx.Header("Access-Control-Allow-Origin", origin)
				ctx.Header("Vary", "Origin")
				ctx.Header("Access-Control-Allow-Credentials", "true")

			}

			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept, Origin, X-Requested-With")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length")

		}

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}

func parseAllowedOrigins(env string) map[string]struct{} {
	result := make(map[string]struct{})
	env = strings.TrimSpace(env)
	if env == "" {
		return result
	}

	for _, raw := range strings.Split(env, ",") {
		origin := strings.TrimSpace(raw)
		if origin == "" {
			continue
		}
		result[origin] = struct{}{}
	}

	return result
}

func isOriginAllowed(origin string, allowed map[string]struct{}) bool {

	if len(allowed) == 0 {
		return true
	}
	_, ok := allowed[origin]
	return ok
}
