package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IssueAuthenticationToken(subject string, secret string) (string, error) {
	now := time.Now().UTC()

	claims := jwt.MapClaims{
		"sub":  subject,
		"iat":  now.UTC(),
		"exp":  now.Add(1 * time.Hour).UTC().Unix(),
		"type": "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedAccessToken, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedAccessToken, nil
}
