package router

import (
	"fmt"
	"net/http"
	handler "url-shortening-service/cmd/api/Handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(addr string, handler *handler.Handler) {
	r := gin.Default()

	r.GET("/:code", func(ctx *gin.Context) {
		handler.Redirect(ctx)
	})

	api := r.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Holaaaa",
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

	fmt.Printf("Server run on: %s", addr)

	r.Run(addr)
}
