package crawler

import (
	"github.com/forestyc/playground/cmd/crawler/app/context"
)

type Crawler interface {
	Init(ctx context.Context, task string)
	Run()
}
