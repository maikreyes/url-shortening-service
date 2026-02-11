package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PutData(ctx *gin.Context) {
	code := ctx.Param("code")
	url := ctx.GetHeader("url")

	newCode, err := h.Service.UpdateShortUrl(code, url)

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
