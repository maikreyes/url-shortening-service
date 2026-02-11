package ports

import (
	"url-shortening-service/internal/domain"
)

type GithubService interface {
	SendMessage(event, code string, payload domain.GithubPayload) (domain.WebhookSend, error)
}
