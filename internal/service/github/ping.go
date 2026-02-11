package github

import (
	"url-shortening-service/internal/domain"
)

func (s *Service) ping(githubPayload domain.GithubPayload) domain.Embed {
	message := "Pong! This is a response to your ping event. for the repository: " + githubPayload.Repository.Name + "\nBy this user: " + githubPayload.Sender.Login

	return domain.Embed{
		Title:       "Ping Response",
		Description: message,
		Color:       0xFFDE21,
	}
}
