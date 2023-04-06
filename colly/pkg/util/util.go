package util

import "strings"

// ReplaceNBSPinHtml 将html中转义后的non-breaking space符号，替换为"&nbsp;"字符串保存（管理端插件kindeditor对转以后\u00a0的格式化有bug）
func ReplaceNBSPinHtml(s string) string {
	return strings.Replace(s, "\u00a0", "&nbsp;", -1)
}
