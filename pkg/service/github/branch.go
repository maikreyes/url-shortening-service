package github

import (
	"url-shortening-service/pkg/domain"
)

func (s *Service) branch(githubPayload domain.GithubPayload) domain.Embed {

	message := "New branch created by: " + githubPayload.Sender.Login + "\nwith this name: " + githubPayload.Ref

	return domain.Embed{
		Title:       "Branch Created",
		Description: message,
		Color:       0x3357FF,
	}
}
