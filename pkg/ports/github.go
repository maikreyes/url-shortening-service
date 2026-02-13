package ports

import (
	"url-shortening-service/pkg/domain"
)

type GithubService interface {
	SendMessage(event, avatarUrl, code string, payload domain.GithubPayload) (domain.WebhookSend, error)
}
