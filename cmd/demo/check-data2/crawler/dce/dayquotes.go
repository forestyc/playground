package dce

import (
	"fmt"
	"github.com/forestyc/playground/cmd/demo/check-data2/model"
	"github.com/forestyc/playground/cmd/demo/check-data2/util"
	"github.com/forestyc/playground/pkg/crawler"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
	"time"
)

type Cache map[string]model.DayQuotes
type SeriesCache map[string][]string

type DayQuotes struct {
	c          *crawler.Colly
	quotType   int
	tradeDate  string
	cache      Cache
	seresCache SeriesCache
}

func NewDayQuotes(date, layout, quotType string) *DayQuotes {
	dq := DayQuotes{
		cache:      make(Cache),
		seresCache: make(SeriesCache),
	}
	dq.quotType, _ = strconv.Atoi(quotType)
	tradeDate, _ := time.ParseInLocation(layout, date, time.Local)
	if quotType == "1" {
		quotType = "0"
	} else {
		quotType = "1"
	}
	postForm := map[string]string{
		"dayQuotes.variety":    "all",
		"dayQuotes.trade_type": quotType,
		"year":                 strconv.Itoa(tradeDate.Year()),
		"month":                strconv.Itoa(int(tradeDate.Month() - 1)),
		"day":                  strconv.Itoa(tradeDate.Day()),
	}

	dq.c = crawler.NewColly(
		"dce-dayquotes",
		"http://www.dce.com.cn/publicweb/quotesdata/dayQuotesCh.html",
		crawler.WithReqType(crawler.Post),
		crawler.WithPostForm(postForm),
		crawler.WithCrawlCallback(dq.Callback()),
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
		dq.c.Crawler.OnHTML(`#printData`, func(element *colly.HTMLElement) {
			element.ForEach(`div`, func(i int, div *colly.HTMLElement) {
				if i == 0 {
					div.ForEach(`tr`, func(i int, tr *colly.HTMLElement) {
						if i != 0 {
							dayQuotes := model.DayQuotes{QuotType: dq.quotType}
							tr.ForEach(`td`, func(j int, td *colly.HTMLElement) {
								if dq.quotType == util.Future {
									dq.parseFuture(&dayQuotes, j, td)
								} else {
									dq.parseOption(&dayQuotes, j, td)
								}
							})
							if dayQuotes.ContractId != "" {
								dq.cache[dayQuotes.ContractId] = dayQuotes
								if dq.quotType == util.Option {
									series := util.GetSeriesFromContract(dayQuotes.ContractId)
									contracts := dq.seresCache[series]
									contracts = append(contracts, dayQuotes.ContractId)
								}
							}
						}
					})
				} else {
					// 补充期权隐波率
					div.ForEach(`tr`, func(i int, tr *colly.HTMLElement) {
						if i != 0 {
							var series, impliedVolatility string
							tr.ForEach(`td`, func(i int, td *colly.HTMLElement) {
								if i == 0 {
									series = strings.TrimSpace(td.Text)
								} else if i == 1 {
									impliedVolatility = strings.TrimSpace(td.Text)
								}
							})
							contracts := dq.seresCache[series]
							for _, contract := range contracts {
								dayQuotes := dq.cache[contract]
								dayQuotes.ImpVol = impliedVolatility
								dq.cache[contract] = dayQuotes
							}
						}
					})
				}
			})
		})
	}
}

func (dp *DayQuotes) parseFuture(dayQuotes *model.DayQuotes, i int, td *colly.HTMLElement) {
	switch i {
	case 1:
		dayQuotes.ContractId = strings.TrimSpace(td.Text)
		dayQuotes.VarietyId = util.GetVarietyFromContract(util.Dce, dayQuotes.ContractId)
	case 2:
		open := strings.TrimSpace(td.Text)
		if open != "-" {
			dayQuotes.OpenPrice = open
		}
	case 3:
		high := strings.TrimSpace(td.Text)
		if high != "-" {
			dayQuotes.HighPrice = high
		}
	case 4:
		low := strings.TrimSpace(td.Text)
		if low != "-" {
			dayQuotes.LowPrice = low
		}
	case 5:
		closeP := strings.TrimSpace(td.Text)
		if closeP != "-" {
			dayQuotes.ClosePrice = closeP
		}
	case 6:
		preSettle := strings.TrimSpace(td.Text)
		if preSettle != "-" {
			dayQuotes.LastSettlePrice = preSettle
		}
	case 7:
		settle := strings.TrimSpace(td.Text)
		if settle != "-" {
			dayQuotes.SettlePrice = settle
		}
	case 10:
		dayQuotes.TotalMatQty, _ = strconv.ParseUint(strings.TrimSpace(td.Text), 10, 64)
	case 11:
		dayQuotes.TotalPos, _ = strconv.ParseUint(strings.TrimSpace(td.Text), 10, 64)
	case 13:
		trunover, err := strconv.ParseFloat(strings.TrimSpace(td.Text), 64)
		if err != nil && !util.InvalidFloat(trunover) {
			dayQuotes.Turnover = util.Round(trunover, 2)
		}
	}
}

func (dp *DayQuotes) parseOption(dayQuotes *model.DayQuotes, i int, td *colly.HTMLElement) {
	switch i {
	case 1:
		dayQuotes.ContractId = strings.TrimSpace(td.Text)
		dayQuotes.VarietyId = util.GetVarietyFromContract(util.Dce, dayQuotes.ContractId)
		dayQuotes.CpFlag = util.GetOptionTypeFromContract(dayQuotes.ContractId)
	case 2:
		open := strings.TrimSpace(td.Text)
		if open != "-" {
			dayQuotes.OpenPrice = open
		}
	case 3:
		high := strings.TrimSpace(td.Text)
		if high != "-" {
			dayQuotes.HighPrice = high
		}
	case 4:
		low := strings.TrimSpace(td.Text)
		if low != "-" {
			dayQuotes.LowPrice = low
		}
	case 5:
		closeP := strings.TrimSpace(td.Text)
		if closeP != "-" {
			dayQuotes.ClosePrice = closeP
		}
	case 6:
		preSettle := strings.TrimSpace(td.Text)
		if preSettle != "-" {
			dayQuotes.LastSettlePrice = preSettle
		}
	case 7:
		settle := strings.TrimSpace(td.Text)
		if settle != "-" {
			dayQuotes.SettlePrice = settle
		}
	case 10:
		dayQuotes.Delta = strings.TrimSpace(td.Text)
	case 11:
		dayQuotes.TotalMatQty, _ = strconv.ParseUint(strings.TrimSpace(td.Text), 10, 64)
	case 12:
		dayQuotes.TotalPos, _ = strconv.ParseUint(strings.TrimSpace(td.Text), 10, 64)
	case 14:
		trunover, err := strconv.ParseFloat(strings.TrimSpace(td.Text), 64)
		if err != nil && !util.InvalidFloat(trunover) {
			dayQuotes.Turnover = trunover
		}
	}
}
