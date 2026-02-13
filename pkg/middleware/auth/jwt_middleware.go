package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		autoHeader := ctx.GetHeader("Authorization")

		if autoHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			ctx.Abort()
			return
		}

		tokestring := strings.TrimPrefix(autoHeader, "Bearer ")

		err := ValidateToken(tokestring)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
