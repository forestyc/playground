package common

const (
	Exchange                  = 52 // 广期所
	Host                      = "http://www.gfex.com.cn"
	Origin                    = "广期所"
	Column                    = "1,4-1"                              // 栏目：市场要闻、交易所信息-最新动态
	PageTotalPrefix           = "createPageHTML('page_div',"         // 页码总数前缀
	PageTotalSuffix           = ", 1,'list_yw','shtml'"              // 页码总数后缀
	NewsPageUrlFirst          = Host + "/gfex/bsyw/list_yw.shtml"    // 本所要闻首页
	NewsPageUrlFormat         = Host + "/gfex/bsyw/list_yw_%d.shtml" // 本所要闻非首页格式
	AnnouncementPageUrlFirst  = Host + "/gfex/tzts/list_yw.shtml"    // 通知公告首页
	AnnouncementPageUrlFormat = Host + "/gfex/tzts/list_yw_%d.shtml" // 通知公告非首页格式
	FocusPageUrlFirst         = Host + "/gfex/mtjj/list_yw.shtml"    // 通知公告首页
	FocusPageUrlFormat        = Host + "/gfex/mtjj/list_yw_%d.shtml" // 通知公告非首页格式
)
