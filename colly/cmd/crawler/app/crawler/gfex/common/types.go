package common

const (
	Host            = "http://www.gfex.com.cn"
	PageTotalPrefix = "createPageHTML('page_div'," // 页码总数前缀
	PageTotalSuffix = ", 1,'list_yw','shtml',25);" // 页码总数猴嘴
	PageUrlFormat   = Host + "/gfex/bsyw/list_yw_%d.shtml"
	Origin          = "广期所"
	Column          = "1,4-1" // 栏目：市场要闻、交易所信息-最新动态
)
