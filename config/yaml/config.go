package yaml

type Config struct {
	//ServerPort  string       `yaml:"serverPort"`
	Db  Db                   `yaml:"db"`
	Rpc  Rpc                 `yaml:"rpc"`
	Env string               `yaml:"env" default:"development"`
	Feign      Feign 		 `yaml:"feign"`
	Application  Application `yaml:"application"`
}

type Db struct{
	Enabled bool  `yaml:"enabled" default:"false"`
	Dsn string      `yaml:"dsn"`
}

type Rpc struct{
	Server Server `yaml:"server"`
}

type Server struct {
	Enabled bool `yaml:"enabled" default:"false"`
	Prefix string `yaml:"prefix"`
}

type Feign struct {
	Service  []Service `yaml:"service"`
	Debug    string `yaml:"debug" default:"none"`
}

type Service struct {
	Name           string `yaml:"name"`
	Url            string `yaml:"url"`
	Enabled        string `yaml:"enabled" default:"true"`
	Path       	   string `yaml:"path"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	TokenKey       string `yaml:"token_key" default:"token"`
	TokenHeader    string `yaml:"token_header" default:"X-Authorization:Bearer"`
	Interval       string `yaml:"interval" default:"3600"`
}

type Application struct {
	ServerPort string       `yaml:"serverPort" default:"9080"`
	JwtSignatureKey string  `yaml:"jwtSignatureKey"`
	JwtExpireIn  string      `yaml:"jwtExpireIn" default:"7200"`
	StaticPath  string       `yaml:"staticPath"`
}
