package config

type Db struct {
	Enabled bool `value:"false"`
	Dsn string
}
