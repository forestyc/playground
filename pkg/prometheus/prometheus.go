package prometheus

type Config struct {
	Addr string `mapstructure:"addr"`
	Path string `mapstructure:"path"`
}
