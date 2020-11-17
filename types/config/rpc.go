package config

import "github.com/duanchi/heron/types/config/rpc"

type Rpc struct {
	Server rpc.Server `yaml:"server"`
}
