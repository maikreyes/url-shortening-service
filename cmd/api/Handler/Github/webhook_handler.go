package github

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
	"url-shortening-service/pkg/domain"

	"github.com/gin-gonic/gin"
)

func (h *Hanlder) WebHookHandler(ctx *gin.Context) {

	code := ctx.Param("code")
	avatarUrl := strings.TrimSpace(ctx.GetHeader("avatarUrl"))
	event := strings.TrimSpace(ctx.GetHeader("X-GitHub-Event"))

	if avatarUrl != "" {
		avatarUrl = "https://ysqz0oydi7thsqmt.public.blob.vercel-storage.com/moik%202.png"
	}

	body, err := ctx.GetRawData()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to read request body",
		})
		return
	}

	var GithubPayload domain.GithubPayload

	if err = json.Unmarshal(body, &GithubPayload); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	send, err := h.Service.SendMessage(event, avatarUrl, code, GithubPayload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if strings.TrimSpace(send.Url) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing discord webhook url",
		})
		return
	}

	b, err := json.Marshal(send.Payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to marshal payload",
		})
		return
	}

	req, err := http.NewRequest(http.MethodPost, send.Url, bytes.NewBuffer(b))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create request",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to send message to discord",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error":  "discord webhook returned non-2xx",
			"status": resp.Status,
			"body":   string(bodyBytes),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "sent",
	})
}
