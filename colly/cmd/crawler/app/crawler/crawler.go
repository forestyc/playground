package crawler

import (
	"github.com/Baal19905/playground/colly/cmd/crawler/app/context"
	"github.com/Baal19905/playground/colly/pkg/log/zap"
)

type Crawler interface {
	Init(ctx context.GlobalContext, config zap.Config)
	Run()
}
