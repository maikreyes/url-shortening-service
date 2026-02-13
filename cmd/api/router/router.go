package router

import (
	"fmt"
	"net/http"
	github "url-shortening-service/cmd/api/Handler/Github"
	handler "url-shortening-service/cmd/api/Handler/urls"
	user "url-shortening-service/cmd/api/Handler/user"
	"url-shortening-service/pkg/middleware/auth"

	"github.com/gin-gonic/gin"
)

func NewRouter(addr string, handler *handler.Handler, GithubHandler *github.Hanlder, userHanlder *user.Handler) {
	r := BuildRouter(handler, GithubHandler, userHanlder)

	fmt.Printf("Server run on: %s", addr)

	_ = r.Run(addr)
}

func BuildRouter(handler *handler.Handler, GithubHandler *github.Hanlder, userHandler *user.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Bienvenido",
		})
	})

	r.GET("/:code", func(ctx *gin.Context) {
		handler.Redirect(ctx)
	})
	r.POST(("/:code/webhook"), func(ctx *gin.Context) {
		GithubHandler.WebHookHandler(ctx)
	})

	r.POST("/login", func(ctx *gin.Context) {
		userHandler.Login(ctx)
	})

	r.POST("/register", func(ctx *gin.Context) {
		userHandler.Register(ctx)
	})

	api := r.Group("/api")
	api.Use(auth.JWTMiddleware())
	{
		v1 := api.Group("/v1")
		v2 := api.Group("/v2")

		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Welcome to triggerito API",
			})
		})

		v1.GET("/shorten/:code", func(ctx *gin.Context) {
			handler.GetData(ctx)
		})
		v1.POST("/shorten", func(ctx *gin.Context) {
			handler.PostData(ctx)
		})
		v1.PUT("/shorten/:code", func(ctx *gin.Context) {
			handler.PutData(ctx)
		})
		v1.DELETE("/shorten/:code", func(ctx *gin.Context) {
			handler.DeleteData(ctx)
		})

		v2.GET("/shorten/:code/stats", func(ctx *gin.Context) {
			handler.GetStats(ctx)
		})

	}

	return r
}
