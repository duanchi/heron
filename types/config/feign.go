package config

import "go.heurd.com/heron-go/heron/types/config/feign"

type Feign struct {
	Enabled bool `yaml:"enabled" default:"false"`
	Services  []feign.Service `yaml:"services"`
	Debug    string `yaml:"debug" default:"false"`
}
