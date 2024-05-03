package main

import (
	"flag"
	"github.com/forestyc/playground/pkg/core/log/zap"
	"github.com/forestyc/playground/pkg/core/redis"
)

var (
	r *redis.Redis

	zlog *zap.Zap
)

func main() {
	var err error
	var operation, keys, dumpFile, address, password string
	flag.StringVar(&operation, "operation", "", "import/export")
	flag.StringVar(&keys, "keys", "", "key1,key2,key3")
	flag.StringVar(&dumpFile, "dump-file", "redis.dump", "redis.dump")
	flag.StringVar(&address, "address", "localhost:6379", "redis address")
	flag.StringVar(&password, "password", "", "redis password")
	flag.Parse()
	r, err = redis.NewRedis(redis.Config{
		Address:          address,
		Password:         password,
		MaxOpen:          1,
		IdleConns:        1,
		IdleTimout:       360,
		OperationTimeout: 60,
	})
	zlog = zap.NewZap(zap.Config{
		Level:         "info",
		Format:        "console",
		Director:      "./log",
		ShowLine:      true,
		EncodeLevel:   "LowercaseColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	})
	if err != nil {
		panic(err)
	}
	if operation == "export" {
		e := NewExport(keys, dumpFile)
		defer e.Close()
		e.Run()
	} else {
		i := NewImport(dumpFile)
		i.Run()
	}
}
