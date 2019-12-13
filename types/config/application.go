package config

type Application struct {
	ServerPort string `yaml:"server_port" env:"SERVER_PORT"`
	StaticPath string `yaml:"static_path" value:""`
}
