package auth

import (
	"github.com/snowmerak/Guards/lib/auth/request"
	"github.com/snowmerak/Guards/lib/auth/script"
)

type Config struct {
	Routes  map[string]request.Config `yaml:"routes"`
	Scripts []script.Config           `yaml:"scripts"`
}
