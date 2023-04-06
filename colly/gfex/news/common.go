package gfex

import "github.com/Baal19905/playground/colly/pkg/crawler"

const (
	host            = "http://www.gfex.com.cn"
	pageTotalPrefix = "createPageHTML('page_div'," // 页码总数前缀
	pageTotalSuffix = ", 1,'list_yw','shtml',25);" // 页码总数猴嘴
	pageUrlFormat   = host + "/gfex/bsyw/list_yw_%d.shtml"
	origin          = "广期所"
	column          = "1,4-1" // 栏目：市场要闻、交易所信息-最新动态
)

// Article 文章
type Article struct {
	PublishDate string
	SortDate    string
	Title       string
	Origin      string
	Body        string
	ColumnLevel string
	RefColumns  string
	crawler     crawler.Colly
}
