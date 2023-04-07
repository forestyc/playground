package crawler

import "github.com/Baal19905/playground/colly/cmd/crawler/app/context"

type Crawler interface {
	Init(ctx context.GlobalContext)
	Run()
}
