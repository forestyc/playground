package main

import (
	"fmt"
	"github.com/forestyc/playground/cmd/demo/gorm/model"
	"github.com/forestyc/playground/pkg/db"
	"time"
)

func main() {
	mdb := db.NewMysql(db.Config{
		Dsn:              "baal:Baal@123@tcp(140.143.163.171:3306)/baal?charset=utf8mb4&parseTime=true&loc=Local",
		MaxOpen:          10,
		IdleConns:        5,
		MaxIdleTime:      300,
		OperationTimeout: 10,
	})
	defer mdb.Close()
	var dq model.DayQuote
	session := mdb.Session()
	tx := session.Begin()
	result := tx.Model(&model.DayQuote{}).
		Take(&dq)
	if code, msg := mdb.DBError(result.Error); code > 0 {
		tx.Rollback()
		fmt.Println(msg)
	}
	dq2 := dq
	dq2.Date = time.Now().AddDate(0, 0, 1)
	result = tx.Model(&model.DayQuote{}).
		Select("date").
		Where("date=? and exchange=? and contract_id=?", dq.Date, dq.Exchange, dq.ContractID).
		Updates(&dq2)
	if code, msg := mdb.DBError(result.Error); code > 0 {
		tx.Rollback()
		fmt.Println(msg)
	}

	result = tx.Model(&model.MarketInfo{}).
		Create(&model.MarketInfo{
			TradeDate: time.Now().AddDate(0, 0, 1).Format("20060102"),
		})
	if code, msg := mdb.DBError(result.Error); code > 0 {
		tx.Rollback()
		fmt.Println(msg)
	}
	tx.Commit()
}
