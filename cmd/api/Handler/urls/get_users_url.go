package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserUrls godoc
// @Summary Listar URLs del usuario
// @Description Lista URLs del usuario (username en path debe coincidir con el del token).
// @Tags urls
// @Produce json
// @Param Authorization header string false "Bearer <token>"
// @Param username path string true "Nombre de usuario"
// @Success 200 {array} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v3/{username}/urls [get]
func (h *Handler) GetUserUrls(ctx *gin.Context) {
	username := ctx.Param("username")

	userVal, exist := ctx.Get("email")

	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unautorized request",
		})
		return
	}

	u, err := h.UserService.GetUserInformation(userVal.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not retrieve user information",
		})
		return
	}

	if u.Username != username {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "forbidden request",
		})
		return
	}

	urls, err := h.UrlService.GetUserUrls(u.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not retrieve urls",
		})
		return
	}

	ctx.JSON(http.StatusOK, urls)
}
