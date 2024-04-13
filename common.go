package auth

import "errors"

var (
	ErrorPassword    = errors.New("invalid password")
	ErrorToken       = errors.New("invalid token")
	ErrorTokenExpiry = errors.New("token expired")
)
	