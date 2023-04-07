package common

import (
	goContext "context"
	"github.com/Baal19905/playground/colly/cmd/crawler/app/context"
	"github.com/go-redis/redis/v8"
)

// Record 记录已爬取网页
func Record(ctx context.GlobalContext, sign string) {
	if ctx.C.Crawler.MarkEnable {
		ctx.Cache.HSet(goContext.Background(), ctx.C.Crawler.Mark, sign, "1")
	}
}

// Skip 判断是否跳过网页
func Skip(ctx context.GlobalContext, sign string) bool {
	if ctx.C.Crawler.MarkEnable {
		result := ctx.Cache.HGet(goContext.Background(), ctx.C.Crawler.Mark, sign)
		if result.Err() != redis.Nil {
			return true
		}
	}
	return false
}
