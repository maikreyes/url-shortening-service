package handler

import (
	"net/http"
	"url-shortening-service/pkg/domain"

	"github.com/gin-gonic/gin"
)

// GetData godoc
// @Summary Obtener short URL
// @Description Obtiene el recurso asociado a un código (requiere auth).
// @Tags urls
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param code path string true "Código corto"
// @Success 200 {object} domain.ApiResponseWithotStats
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/shorten/{code} [get]
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
