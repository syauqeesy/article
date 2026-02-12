package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func IssueAuthenticationToken(subject string, ttl time.Duration, secret string) (*string, *int64, error) {
	now := time.Now().UTC()

	expiresAt := now.Add(ttl).UTC().UnixMilli()

	claims := jwt.MapClaims{
		"sub":  subject,
		"iat":  now.UTC(),
		"exp":  now.Add(ttl).UTC().Unix(),
		"type": "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedAccessToken, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return nil, nil, err
	}

	return &signedAccessToken, &expiresAt, nil
}

