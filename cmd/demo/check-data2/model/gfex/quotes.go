package gfex

type QuotesResponse struct {
	Code  string `json:"code"`
	Msg   string `json:"msg"`
	Param struct {
		TradeDate []string `json:"trade_date"`
		TradeType []string `json:"trade_type"`
	} `json:"param"`
	Data []struct {
		Variety           string   `json:"variety"`
		DiffI             *int     `json:"diffI"`
		High              *float64 `json:"high"`
		Turnover          *float64 `json:"turnover"`
		ImpliedVolatility *float64 `json:"impliedVolatility"`
		Diff              *int     `json:"diff"`
		Delta             *float64 `json:"delta"`
		Close             *float64 `json:"close"`
		Diff1             *float64 `json:"diff1"`
		LastClear         *float64 `json:"lastClear"`
		Open              *float64 `json:"open"`
		MatchQtySum       *int     `json:"matchQtySum"`
		DelivMonth        string   `json:"delivMonth"`
		Low               *float64 `json:"low"`
		ClearPrice        *float64 `json:"clearPrice"`
		VarietyOrder      string   `json:"varietyOrder"`
		OpenInterest      *int     `json:"openInterest"`
		Volumn            *int     `json:"volumn"`
	} `json:"data"`
	Time int64 `json:"time"`
}
