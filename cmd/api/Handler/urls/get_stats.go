package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetStats(ctx *gin.Context) {

	code := ctx.Param("code")
	data, err := h.Service.GetShortUrl(code)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "code not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
