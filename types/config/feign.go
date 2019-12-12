package config

import "go.heurd.com/heron-go/heron/types/config/feign"

type Feign struct {
	Services  []feign.Service `yaml:"service"`
	Debug    string `yaml:"debug" default:"none"`
}
