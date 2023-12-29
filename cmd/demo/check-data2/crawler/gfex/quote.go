package gfex

import (
	"encoding/json"
	"fmt"
	"github.com/forestyc/playground/cmd/demo/check-data2/model"
	"github.com/forestyc/playground/cmd/demo/check-data2/model/gfex"
	"github.com/forestyc/playground/cmd/demo/check-data2/util"
	"github.com/forestyc/playground/pkg/crawler"
	"github.com/forestyc/playground/pkg/db"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
	"time"
)

type DayQuotes struct {
	c         *crawler.Colly
	mydb      *db.Mysql
	quotType  int
	tradeDate string
	response  gfex.QuotesResponse
	cache     []model.DayQuotes
}

func NewDayQuotes(date, layout, quotType string) *DayQuotes {
	dq := DayQuotes{
		mydb: db.NewMysql(db.Config{
			Dsn:              "baal:Baal@123@tcp(140.143.163.171:3306)/baal?charset=utf8",
			MaxOpen:          10,
			IdleConns:        10,
			MaxIdleTime:      300,
			OperationTimeout: 60,
		}),
	}
	dq.quotType, _ = strconv.Atoi(quotType)
	tradeDate, _ := time.ParseInLocation(layout, date, time.Local)
	if quotType == "1" {
		quotType = "0"
	} else {
		quotType = "1"
	}
	dq.c = crawler.NewColly(
		"dce-dayquotes",
		"http://www.gfex.com.cn/u/interfacesWebTiDayQuotes/loadList?trade_date="+tradeDate.Format("20060102")+"&trade_type="+quotType,
		crawler.WithReqType(crawler.Post),
		crawler.WithCrawlCallback(dq.Callback()),
		crawler.WithPipeline(dq.Pipeline()),
	)
	return &dq
}

func (dq *DayQuotes) Run() {
	dq.c.Run()
}

func (dq *DayQuotes) Print() {
	fmt.Println(dq.cache, len(dq.cache))
}

func (dq *DayQuotes) Callback() crawler.Callback {
	return func() {
		dq.c.Crawler.OnResponse(func(response *colly.Response) {
			if err := json.Unmarshal(response.Body, &dq.response); err != nil {
				fmt.Println(err)
				return
			}
			if dq.response.Code == "0" {
				for _, item := range dq.response.Data {
					if strings.Contains(item.VarietyOrder, "计") ||
						strings.TrimSpace(item.VarietyOrder) == "" ||
						strings.TrimSpace(item.DelivMonth) == "" {
						continue
					}
					quotes := model.DayQuotes{}
					quotes.Exchange = "GFEX"
					quotes.VarietyId = item.VarietyOrder
					quotes.QuotType = dq.quotType
					if quotes.QuotType == util.Option {
						quotes.ContractId = item.DelivMonth
					} else if quotes.QuotType == util.Future {
						quotes.ContractId = item.VarietyOrder + item.DelivMonth
					}
					if item.Open != nil {
						quotes.OpenPrice = fmt.Sprintf("%.2f", util.Round(*item.Open, 2))
					}
					if item.Close != nil {
						quotes.ClosePrice = fmt.Sprintf("%.2f", util.Round(*item.Close, 2))
					}
					if item.High != nil {
						quotes.HighPrice = fmt.Sprintf("%.2f", util.Round(*item.High, 2))
					}
					if item.Low != nil {
						quotes.LowPrice = fmt.Sprintf("%.2f", util.Round(*item.Low, 2))
					}
					if item.LastClear != nil {
						quotes.SettlePrice = fmt.Sprintf("%.2f", util.Round(*item.LastClear, 2))
					}
					if item.ClearPrice != nil {
						quotes.LastSettlePrice = fmt.Sprintf("%.2f", util.Round(*item.ClearPrice, 2))
					}
					if item.Volumn != nil {
						quotes.TotalMatQty = uint64(*item.Volumn)
					}
					if item.OpenInterest != nil {
						quotes.TotalMatQty = uint64(*item.OpenInterest)
					}
					if item.Turnover != nil {
						quotes.Turnover = util.Round(*item.Turnover, 2)
					}
					if item.Delta != nil {
						quotes.Delta = fmt.Sprintf("%.2f", util.Round(*item.Delta, 2))
					}
					if item.ImpliedVolatility != nil {
						quotes.ImpVol = fmt.Sprintf("%.2f", util.Round(*item.ImpliedVolatility, 2))
					}
					quotes.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
					quotes.CpFlag = util.GetOptionTypeFromContract(quotes.ContractId)
					quotes.Date = time.Now().Format("20060102")
					dq.cache = append(dq.cache, quotes)
				}
			} else {
				fmt.Println(dq.response.Msg)
			}
		})
	}
}

func (dq *DayQuotes) Pipeline() crawler.Pipeline {
	return func() error {
		session, _ := dq.mydb.Session()
		session = session.Begin()
		session = session.Table("day_quotes").Save(&dq.cache)
		if session.Error != nil {
			fmt.Println("session.Error")
			return session.Error
		}
		session.Commit()
		return nil
	}
}
