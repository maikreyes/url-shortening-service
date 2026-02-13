package handler

import (
	"net/http"
	"url-shortening-service/pkg/domain"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetData(ctx *gin.Context) {
	code := ctx.Param("code")
	usernameVal, exist := ctx.Get("email")

	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unautorized request",
		})
		return
	}

	data, err := h.UrlService.GetShortUrl(code)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "code not found",
		})
		return
	}
	username := usernameVal.(string)

	u, err := h.UserService.GetUserInformation(username)

	if data.UserID != u.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "this user don't have this url",
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
