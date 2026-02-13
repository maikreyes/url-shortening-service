package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.TrimSpace(ctx.GetHeader("Authorization"))
		if tokenString == "" {
			cookieToken, err := ctx.Cookie("access_token")
			if err == nil {
				tokenString = strings.TrimSpace(cookieToken)
			}
		}

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token required",
			})
			ctx.Abort()
			return
		}

		tokenString = strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer "))
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token required",
			})
			ctx.Abort()
			return
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			ctx.Abort()
			return
		}

		emailVal, ok := claims["email"]
		email, okStr := emailVal.(string)
		if !ok || !okStr || strings.TrimSpace(email) == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			ctx.Abort()
			return
		}

		ctx.Set("email", email)

		ctx.Next()
	}
}
