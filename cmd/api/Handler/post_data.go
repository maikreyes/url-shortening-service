package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostData(ctx *gin.Context) {
	rawURL := strings.TrimSpace(ctx.GetHeader("url"))
	if rawURL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "url is required",
		})
		return
	}

	code, err := h.Service.CreateShortUrl(rawURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not create short url",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"url": code,
	})
}
