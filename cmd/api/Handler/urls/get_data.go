package handler

import (
	"net/http"
	"url-shortening-service/internal/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetData(ctx *gin.Context) {
	code := ctx.Param("code")
	data, err := h.Service.GetShortUrl(code)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "code not found",
		})
		return
	}

	dataWithoutStats := domain.ApiResponseWithotStats{
		ID:        data.ID,
		Url:       data.Url,
		ShortCode: data.ShortCode,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, dataWithoutStats)
}
