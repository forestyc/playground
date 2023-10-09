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
		if result.Err() != redis.Nil {
			return true
		}
	}
	return false
}

// GetVariety 获取品种
func GetVariety(ctx context.Context, exchange []int) ([]string, error) {
	if len(exchange) == 0 {
		return nil, errors.New("invalid exchange")
	}
	var variety []string
	session, cancel := ctx.Db.Session()
	defer cancel()
	if err := session.Table("variety").
		Select("variety_id").
		Where("variety_type = 1 and exchange in ?", exchange).
		Find(&variety).Error; err != nil {
		return nil, err
	}
	return variety, nil
}
