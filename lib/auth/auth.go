package auth

type Config struct {
	Type string `yaml:"type"`
	Key  string `yaml:"key"`
}

type Auth interface {
	Authenticate(request Request) error
}

var (
	ErrInvalidCredentials = "invalid credentials"
	ErrInvalidToken       = "invalid token"
	ErrInvalidApiKey      = "invalid api key"
)
