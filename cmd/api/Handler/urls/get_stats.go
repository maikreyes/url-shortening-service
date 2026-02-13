package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStats godoc
// @Summary Obtener estadísticas
// @Description Retorna el recurso con `accessCount`.
// @Tags stats
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param code path string true "Código corto"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v2/shorten/{code}/stats [get]
func (h *Handler) GetStats(ctx *gin.Context) {
	code := ctx.Param("code")
	data, err := h.UrlService.GetShortUrl(code)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "code not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
