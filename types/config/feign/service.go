package feign

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
