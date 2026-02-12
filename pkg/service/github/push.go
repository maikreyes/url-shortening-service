package github

import "url-shortening-service/pkg/domain"

func (s *Service) push(githubPayload domain.GithubPayload) domain.Embed {

	message := "New push on this brach: " + githubPayload.Ref + "\nBy this user: " + githubPayload.Pusher.Name

	return domain.Embed{
		Title:       "Push Event",
		Description: message,
		Color:       0x33FF57,
	}
}
