package common

import (
	goContext "context"
	"errors"
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/go-redis/redis/v8"
)

// Record 记录已爬取网页
func Record(ctx context.Context, sign string) {
	if ctx.C.Crawler.MarkEnable {
		ctx.Cache.HSet(goContext.Background(), ctx.C.Crawler.Mark, sign, "1")
	}
}

// Skip 判断是否跳过网页
func Skip(ctx context.Context, sign string) bool {
	if ctx.C.Crawler.MarkEnable {
		result := ctx.Cache.HGet(goContext.Background(), ctx.C.Crawler.Mark, sign)
		if !errors.Is(result.Err(), redis.Nil) {
			return true
		}
	}
	return false
}
