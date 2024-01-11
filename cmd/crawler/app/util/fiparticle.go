package util

import (
	"github.com/gocolly/colly/v2"
	"strings"
)

// ReplaceNBSPinHtml 将html中转义后的non-breaking space符号，替换为"&nbsp;"字符串保存（管理端插件kindeditor对转以后\u00a0的格式化有bug）
func ReplaceNBSPinHtml(s string) string {
	return strings.Replace(s, "\u00a0", "&nbsp;", -1)
}

// AddHost 文章中带图片的需要添加host
func AddHost(content, host string) string {
	result := content
	if strings.Contains(content, "<img") {
		var srcs []string
		for {
			// img标签
			imgStart := strings.Index(content, "<img")
			if imgStart == -1 {
				break
			}
			imgLabel := content[imgStart:]
			imgEnd := strings.Index(imgLabel, ">")
			if imgEnd == -1 {
				return content
			}
			imgLabel = imgLabel[0 : imgEnd+1]
			// src属性
			srcStart := strings.Index(imgLabel, `src`)
			if srcStart == -1 {
				return content
			}
			srcAttr := imgLabel[srcStart+3:]
			srcEqual := strings.Index(srcAttr, `=`)
			if srcEqual == -1 {
				return content
			}
			srcAttr = srcAttr[srcEqual+1:]
			srcFirstQuot := strings.Index(srcAttr, `"`)
			if srcFirstQuot == -1 {
				return content
			}
			srcLastQuot := strings.Index(srcAttr[srcFirstQuot+1:], `"`)
			if srcLastQuot == -1 {
				return content
			}
			srcAttr = srcAttr[srcFirstQuot+1 : srcLastQuot+1]
			content = content[imgStart+4:]
			srcs = append(srcs, srcAttr)
		}
		for _, src := range srcs {
			if !strings.Contains(src, host) {
				result = strings.ReplaceAll(result, src, host+src)
			}
		}
	}
	return result
}

func Host(element *colly.HTMLElement) string {
	return element.Request.URL.Scheme + "://" + element.Request.URL.Host
}

func FormatArticle(content string, element *colly.HTMLElement) string {
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\n", "")
	content = strings.ReplaceAll(content, "\r", "")
	return AddHost(content, Host(element))
}
