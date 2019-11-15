package config

import "heurd.com/wand-go/wand/types/config/rpc"

type Rpc struct {
	Server struct{
		rpc.Server
	}
}
