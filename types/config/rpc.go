package config

import "go.heurd.com/heron-go/heron/types/config/rpc"

type Rpc struct {
	Server rpc.Server `yaml:"server"`
}
