//go:build !windows
// +build !windows

package zap

import (
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

func GetWriteSyncer(z Config) (zapcore.WriteSyncer, error) {
	fileWriter, err := zaprotatelogs.New(
		path.Join(z.Director, "%Y-%m-%d.log"),
		zaprotatelogs.WithLinkName(z.LinkName),
		zaprotatelogs.WithMaxAge(60*24*time.Hour), //保留60天
		zaprotatelogs.WithRotationTime(24*time.Hour),
	)
	if z.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
