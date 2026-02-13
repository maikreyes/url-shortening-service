package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
