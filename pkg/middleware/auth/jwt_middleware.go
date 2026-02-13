package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		autoHeader, err := ctx.Cookie("access_token")

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "error to try obtenin cookie",
			})
		}

		if autoHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			ctx.Abort()
			return
		}

		tokestring := strings.TrimPrefix(autoHeader, "Bearer ")

		err = ValidateToken(tokestring)
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
