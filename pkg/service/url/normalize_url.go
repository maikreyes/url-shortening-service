package url

import "strings"

func normalizeURL(raw string) string {
	u := strings.TrimSpace(raw)
	if u == "" {
		return u
	}

	if strings.HasPrefix(u, "//") {
		return "https:" + u
	}

	if strings.Contains(u, "://") {
		return u
	}

	return "https://" + u
}
