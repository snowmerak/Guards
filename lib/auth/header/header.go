package header

import (
	"errors"
)

type Config struct {
	Requires []string `yaml:"requires"`
}

type Headers map[string]string

func (h Headers) Get(key string) string {
	return h[key]
}

var (
	ErrInvalidHeader = errors.New("invalid header")
	ErrMissingHeader = errors.New("missing header")
)
