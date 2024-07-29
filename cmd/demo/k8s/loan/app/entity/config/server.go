package config

type Server struct {
	Addr string `mapstructure:"addr"`
	Id   int    `mapstructure:"id"`
}
