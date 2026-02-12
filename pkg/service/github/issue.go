package github

import (
	"url-shortening-service/pkg/domain"
)

func (s *Service) issue(githubPayload domain.GithubPayload) domain.Embed {

	message := "New issue " + githubPayload.Action + " by " + githubPayload.Sender.Login + " on " + githubPayload.Repository.Name

	var color int

	switch githubPayload.Action {
	case "opened":
		color = 0x33FF57
	case "closed":
		color = 0xC70039
	case "reopened":
		color = 0x900C3F
	default:
		color = 0x581845
	}

	return domain.Embed{
		Title:       "Issue " + githubPayload.Action,
		Description: message,
		Color:       color,
	}
}
