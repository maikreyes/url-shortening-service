package domain

type WebhookSend struct {
	Payload DiscordPayload `json:"payload"`
	Url     string         `json:"url"`
}
