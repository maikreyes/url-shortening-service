package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteData(ctx *gin.Context) {
	code := ctx.Param("code")
	err := h.UrlService.DeleteShortUrl(code)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "code not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "code deleted successfully",
	})
}
