package service

import "time"

type Periods struct {
	periods map[string]int
}

func NewPeriods(startDate string, periods int) *Periods {
	p := &Periods{
		periods: make(map[string]int),
	}

	for i := 0; i < periods; i++ {
		p.AddPeriod(startDate, i)
		t, _ := time.ParseInLocation("2006-01-02", startDate, time.Local)
		t.AddDate(0, 1, 0)
		startDate = t.Format("2006-01-02")
	}
	return p
}

func (p *Periods) AddPeriod(period string, periodIndex int) {
	p.periods[period] = periodIndex
}

func (p *Periods) GetPeriod(period string) int {
	return p.periods[period]
}
