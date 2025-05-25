package auth

import "github.com/snowmerak/Guards/lib/auth/request"

type Auth interface {
	Authenticate(request request.Request) error
}

var (
	ErrInvalidCredentials = "invalid credentials"
	ErrInvalidToken       = "invalid token"
	ErrInvalidApiKey      = "invalid api key"
)
