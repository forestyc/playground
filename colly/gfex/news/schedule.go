package gfex

import "fmt"

// CrawlerNews 广期所-本所要闻
func CrawlerNews() error {
	// 获取文章href
	href := CrawlerArticleHref()
	// 爬取文章
	for _, e := range href {
		a := CrawlerArticle(e)
		fmt.Println(a.Title)
	}
	return nil
}

// CrawlerArticleHref 获取所有文章
func CrawlerArticleHref() []string {
	listPage := NewListPage()
	listPage.Run()
	return listPage.GetArticleHref()
}

func CrawlerArticle(url string) *Article {
	news := NewNews(url)
	news.Run()
	return news
}
