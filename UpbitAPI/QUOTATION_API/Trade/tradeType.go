package Trade

type Trade struct{}

type GetTradeTickInputType struct {
	Market  string
	To      string
	Count   int    // 체결 개수
	Cursor  string // 페이지네이션 커서
	DaysAgo int
}
type GetTradeTickInputParameterType struct {
	Market  string
	To      string
	Count   string
	Cursor  string
	DaysAgo string
}

type TradeTickInfo struct {
	Market           string  `json:"market"`
	TradeDateUTC     string  `json:"trade_date_utc"`
	TradeTimeUTC     string  `json:"trade_time_utc"`
	Timestamp        int     `json:"timestamp"`
	TradePrice       float64 `json:"trade_price"`
	TradeVolume      float64 `json:"trade_volume"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	ChangePrice      float64 `json:"change_price"`
	AskBid           string  `json:"ask_bid"`
	SequentialId     int     `json:"sequential_id"`
}

/*
{
"market":"KRW-ETH",
"trade_date_utc":"2023-02-04",
"trade_time_utc":"09:53:43",
"timestamp":1675504423087,
"trade_price":2090000.00000000,
"trade_volume":0.00747377,
"prev_closing_price":2091000.00000000,
"change_price":-1000.00000000,
"ask_bid":"BID",
"sequential_id":1675504423087000
}
*/
