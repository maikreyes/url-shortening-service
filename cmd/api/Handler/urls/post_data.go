package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostData(ctx *gin.Context) {

	rawURL := strings.TrimSpace(ctx.GetHeader("url"))
	webhookHeader := strings.TrimSpace(ctx.GetHeader("webhook"))

	usernameVal, exist := ctx.Get("email")

	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unautorized request",
		})
		return
	}

	isWebhook := false

	if webhookHeader != "" {
		parsed, err := strconv.ParseBool(webhookHeader)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "webhook header must be a boolean (true/false)",
			})
			return
		}
		isWebhook = parsed
	}

	if rawURL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "url is required",
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

	code, err := h.UrlService.CreateShortUrl(rawURL, u.Username, isWebhook, u.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not create short url",
		})
		return
	}

	if isWebhook {
		code += "/webhook"
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"url": h.Host + "/" + code,
	})
}
