package url

import "strings"

func isDuplicateShortCodeError(err error) bool {
	if err == nil {
		return false
	}

	msg := strings.ToLower(err.Error())

	// MySQL: "Error 1062 ... Duplicate entry ... for key '...idx_...short_code'"
	// Postgres: "duplicate key value violates unique constraint ..." (SQLSTATE 23505)
	if !(strings.Contains(msg, "duplicate") || strings.Contains(msg, "sqlstate 23505")) {
		return false
	}

	return strings.Contains(msg, "short_code") || strings.Contains(msg, "short code") || strings.Contains(msg, "idx")
}
