package model

type DayQuotes struct {
	Exchange        string
	VarietyId       string
	ContractId      string
	OpenPrice       string
	ClosePrice      string
	HighPrice       string
	LowPrice        string
	SettlePrice     string
	LastSettlePrice string
	TotalMatQty     uint64
	TotalPos        uint64
	Turnover        float64
	Delta           string
	ImpVol          string
	CreatedAt       string
	QuotType        int
	CpFlag          string
	Date            string
}
