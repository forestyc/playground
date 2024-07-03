package config

import (
	"github.com/forestyc/playground/pkg/core/db"
	"github.com/forestyc/playground/pkg/core/log/zap"
	"github.com/forestyc/playground/pkg/core/prometheus"
	"github.com/forestyc/playground/pkg/core/redis"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

// Config 配置信息
type Config struct {
	Server     Server            `mapstructure:"server"`
	Database   db.Config         `mapstructure:"database"`
	Log        zap.Config        `mapstructure:"log"`
	Redis      redis.Config      `mapstructure:"redis"`
	Prometheus prometheus.Config `mapstructure:"prometheus"`
}

// Load 加载配置
func Load(file string, c *Config) error {
	// 导入配置文件
	fileName := strings.Split(filepath.Base(file), ".")
	if len(fileName) != 2 {
		return errors.New("invalid config file[" + file + "]")
	}
	dir := filepath.Dir(file)
	viper.SetConfigName(fileName[0])
	viper.SetConfigType(fileName[1])
	viper.AddConfigPath(dir)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err := Unmarshal(c)
		if err != nil {
			panic(err)
		}
	})
	return Unmarshal(c)
}

func Unmarshal(c *Config) error {
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	return nil
}
