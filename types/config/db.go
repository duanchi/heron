package config

type Db struct {
	Enabled bool `yaml:"enabled" default:"false"`
	Dsn string	`yaml:"dsn"`
}
