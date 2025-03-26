package utils

import (
	"errors"
	"strings"
)

func ExtractBearerToken(bearerToken string) (string, error) {
	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 {
		return "", errors.New("invalid bearer token")

	}

	token := parts[1]
	return token, nil
}
