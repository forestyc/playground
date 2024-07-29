package model

import "time"

type RepaymentResp struct {
	Business      float64
	ProvidentFund float64
	Total         float64
}

type GetInfoResp struct {
	BusinessPrincipal         float64
	BusinessInterestRate      float64
	BusinessPeriods           int
	ProvidentFundPrincipal    float64
	ProvidentFundInterestRate float64
	ProvidentPeriods          int
	StartDate                 *time.Time
}
