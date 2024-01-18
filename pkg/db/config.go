package db

type Config struct {
	Dsn              string `mapstructure:"dsn"`
	MaxOpen          int    `mapstructure:"max-open"`
	IdleConns        int    `mapstructure:"idle-conns"`
	MaxIdleTime      int    `mapstructure:"idle-timeout"`
	OperationTimeout int    `mapstructure:"operation-timeout"`
}
