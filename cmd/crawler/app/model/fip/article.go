package fip

const (
	sep  = ","
	tail = "\n"
)

type Article struct {
	Id      int    // 序号
	Title   string // 标题名称
	Content string // 内容正文
	Source  string // 文件来源
	Summary string // 摘要
	Date    string // 发布日期
}
