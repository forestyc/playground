package main

import (
	"encoding/json"
	"fmt"
	"github.com/forestyc/playground/cmd/demo/check-data2/crawler/gfex"
	"github.com/forestyc/playground/cmd/demo/check-data2/model"
	"github.com/forestyc/playground/pkg/http"
	"time"
)

var (
	holidays map[string]struct{}
)

func main() {
	//var date string
	//var quoteType string
	//flag.StringVar(&date, "date", "", "20231229")
	//flag.StringVar(&quoteType, "quote-type", "", "future:1, option:2")
	//flag.Parse()

	dq := gfex.NewDayQuotes("20231229", "20060102", "1")
	//dq := gfex.NewDayQuotes(date, "20060102", quoteType)
	dq.Run()
	return
	//dq := dce.NewDayQuotes("20231227", "20060102", "2")
	//dq.Run()
	//dq.Print()

	//return
	//exchanges := []string{"DCE", "CZCE", "CFFEX", "SHFE", "INE", "GFEX"}
	initHolidays([]string{"20230101", "20230102", "20230121", "20230122", "20230123", "20230124", "20230125", "20230126", "20230127", "20230405", "20230429", "20230430", "20230501", "20230502", "20230503", "20230622", "20230623", "20230624", "20230929", "20230930", "20231001", "20231002", "20231003", "20231004", "20231005", "20231006"})
	dates := getTradeDates("20230301", "20230331")
	fmt.Println(dates)
	//for _, exchange := range exchanges {
	//	fmt.Println("start check", exchange)
	//	checkQuotes(exchange, dates)
	//	fmt.Println("end check", exchange)
	//}
}

type ReqInfo struct {
	Exchange string `json:"exchange"`
	Date     string `json:"date"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	QuotType int    `json:"quotType"`
}

type RspInfo struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Object  struct {
		Quots []model.DayQuotes `json:"quots"`
	} `json:"object"`
}

func checkQuotes(exchange string, dates []string) {
	cli := http.NewClient(true)
	defer cli.Close()
	for _, date := range dates {
		reqFur := ReqInfo{
			Exchange: exchange,
			Date:     date,
			Page:     1,
			PageSize: 10,
			QuotType: 1,
		}
		req, _ := json.Marshal(reqFur)
		rsp, _ := cli.Do("POST", "https://fipinfo.dfitc.com.cn/quot/quot-info", nil, req)
		rspFtr := RspInfo{}
		json.Unmarshal(rsp, &rspFtr)
		if len(rspFtr.Object.Quots) == 0 {
			fmt.Println(exchange, date, "missing ftr data")
		}
		reqOpt := ReqInfo{
			Exchange: exchange,
			Date:     date,
			Page:     1,
			PageSize: 10,
			QuotType: 2,
		}
		req, _ = json.Marshal(reqOpt)
		cli.Do("POST", "https://fipinfo.dfitc.com.cn/quot/quot-info", nil, req)
		rspOpt := RspInfo{}
		json.Unmarshal(rsp, &rspOpt)
		if len(rspOpt.Object.Quots) == 0 {
			fmt.Println(exchange, date, "missing opt data")
		}
	}
}

func getTradeDates(start, end string) []string {
	var dates []string
	timeStart, err := time.ParseInLocation("20060102", start, time.Local)
	if err != nil {
		panic(err)
	}
	timeEnd, err := time.ParseInLocation("20060102", end, time.Local)
	if err != nil {
		panic(err)
	}
	for {
		if timeStart.After(timeEnd) {
			break
		}
		if timeStart.Weekday() != time.Sunday && timeStart.Weekday() != time.Saturday {
			date := timeStart.Format("20060102")
			if _, ok := holidays[date]; !ok {
				dates = append(dates, date)
			}
		}

		timeStart = timeStart.AddDate(0, 0, 1)
	}
	return dates
}

func initHolidays(hs []string) {
	holidays = make(map[string]struct{})
	for _, h := range hs {
		holidays[h] = struct{}{}
	}
}
