package config

type Application struct {
	ServerPort string `env:"SERVER_PORT"`
	StaticPath string `value:""`
}
