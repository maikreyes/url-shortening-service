package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Redirect godoc
// @Summary Redirigir por código
// @Description Redirige (301) al destino asociado al código.
// @Tags redirect
// @Param code path string true "Código corto"
// @Success 301 {string} string "Redirect"
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{code} [get]
func (h *Handler) Redirect(ctx *gin.Context) {
	code := ctx.Param("code")
	data, err := h.UrlService.GetShortUrl(code)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "code not found",
		})
		return
	}

	target := strings.TrimSpace(data.Url)
	if target == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "code not found",
		})
		return
	}

	err = h.UrlService.AddCount(*data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not update count",
		})
		return
	}

	if strings.HasPrefix(target, "//") {
		target = "https:" + target
	} else if !strings.Contains(target, "://") {
		target = "https://" + target
	}

	ctx.Redirect(http.StatusMovedPermanently, target)
}
