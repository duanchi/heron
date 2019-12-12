package config

type Db struct {
	Enabled bool `yaml:"enabled" value:"false"`
	Dsn string	`yaml:"dsn"`
}
