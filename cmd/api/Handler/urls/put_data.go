package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PutData godoc
// @Summary Actualizar short URL
// @Description Actualiza la URL original asociada al código. La nueva URL se envía en el header `url`.
// @Tags urls
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param code path string true "Código corto"
// @Param url header string true "Nueva URL"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/shorten/{code} [put]
func (h *Handler) PutData(ctx *gin.Context) {
	code := ctx.Param("code")
	url := ctx.GetHeader("url")

	usernameVal, exist := ctx.Get("email")

	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unautorized request",
		})
		return
	}

	username := usernameVal.(string)

	u, err := h.UserService.GetUserInformation(username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	newCode, err := h.UrlService.UpdateShortUrl(code, u.Username, url)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not update short url",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Url shorter updated with this new shorted url: " + newCode,
	})

}
