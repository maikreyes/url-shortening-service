package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteData godoc
// @Summary Eliminar short URL
// @Description Elimina el recurso asociado al código.
// @Tags urls
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param code path string true "Código corto"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/shorten/{code} [delete]
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
