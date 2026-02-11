package github

import "url-shortening-service/internal/domain"

func (s *Service) SendMessage(event, code string, githubPayload domain.GithubPayload) (domain.WebhookSend, error) {

	webhook, err := s.Repository.FetchData(code)

	discordPayload := domain.DiscordPayload{
		Username:  "Archives Bot",
		AvatarUrl: "https://avatars.githubusercontent.com/u/191653228?s=96&v=4",
	}

	if err != nil {
		discordPayload.Embeds = []domain.Embed{{
			Title:       "Error",
			Description: "Invalid code provided",
			Color:       16711680,
		}}

		return domain.WebhookSend{
			Payload: discordPayload,
			Url:     "",
		}, err
	}

	var embed domain.Embed

	switch event {
	case "ping":
		embed = s.ping(githubPayload)
	case "issues":
		embed = s.issue(githubPayload)
	case "create":
		embed = s.branch(githubPayload)
	case "push":
		embed = s.push(githubPayload)
	case "pull_request":
		embed = s.pullRequest(githubPayload)
	default:
		embed = domain.Embed{
			Title:       "Unsupported event",
			Description: "Event: " + event,
			Color:       0xFF5733,
		}
	}

	discordPayload.Embeds = []domain.Embed{embed}

	return domain.WebhookSend{
		Payload: discordPayload,
		Url:     webhook.Url,
	}, nil

}
