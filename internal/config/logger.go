package config

type Logger struct {
	Level string `envconfig:"LEVEL" default:"debug"`
}
