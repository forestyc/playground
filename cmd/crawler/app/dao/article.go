package dao

import (
	"github.com/forestyc/playground/cmd/crawler/app/context"
)

type Article struct {
	PublishDate string `gorm:"publish_date"`
	SortDate    string `gorm:"sort_date"`
	Title       string `gorm:"title"`
	Origin      string `gorm:"origin"`
	Body        string `gorm:"body"`
	ColumnLevel string `gorm:"column_level"`
	Species     string `gorm:"species"`
}

// Save 保存
func (a *Article) Create(ctx context.Context) error {
	session, cancel := ctx.Db.Session()
	defer cancel()
	session = session.Begin()
	session = session.Table(ctx.C.Crawler.TableArticles).Create(a)
	if session.Error != nil {
		return session.Error
	}
	session.Commit()
	return nil
}
