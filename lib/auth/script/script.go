package script

type Parameters map[string]any

type Config struct {
	Name   string `yaml:"name"`
	Script string `yaml:"script"`
}

type Script string
