package github

import "url-shortening-service/internal/ports"

type Hanlder struct {
	Service ports.GithubService
}

func NewHandler(service ports.GithubService) *Hanlder {
	return &Hanlder{
		Service: service,
	}
}
