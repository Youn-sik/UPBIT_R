package Candle

type Candle struct{}

type GetMinCandleInputType struct {
	MinUnit int
	Market  string
	To      string
	Count   int
}
type GetMinCandleInputParameterType struct {
	MinUnit string
	Market  string
	To      string
	Count   string
}

type MinCandleInfo struct {
	Market               string  `json:"market"`
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"` // 캔들 기준 시각
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`           // 시가
	HighPrice            float64 `json:"high_price"`              // 고가
	LowPrice             float64 `json:"low_price"`               // 저가
	TradePrice           float64 `json:"trade_price"`             // 종가
	Timestamp            int     `json:"timestamp"`               // 해당 캔들에서 마지막 틱이 지정된 시각
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`  // 누적 거래 금액
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"` // 누적 거래량
	Unit                 int     `json:"unit"`                    // 분 단위
}

/*
{
"market":"KRW-ETH",
"candle_date_time_utc":"2023-02-04T01:33:00",
"candle_date_time_kst":"2023-02-04T10:33:00",
"opening_price":2096000.00000000,
"high_price":2097000.00000000,
"low_price":2096000.00000000,
"trade_price":2097000.00000000,
"timestamp":1675474425744,
"candle_acc_trade_price":25451270.17213000,
"candle_acc_trade_volume":12.13985699,
"unit":1
}
*/

type GetDayCandleInputType struct {
	Market              string
	To                  string
	Count               int
	ConvertingPriceUnit string
}
type GetDayCandleInputParameterType struct {
	Market              string
	To                  string
	Count               string
	ConvertingPriceUnit string
}

type DayCandleInfo struct {
	Market               string  `json:"market"`
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TradePrice           float64 `json:"trade_price"`
	Timestamp            int     `json:"timestamp"`
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
	PrevClosingPrice     float64 `json:"prev_closing_price"`    // 전일 종가
	ChargePrice          float64 `json:"change_price"`          // 전일 종가 대비 변화 금액
	ChargeRate           float64 `json:"change_rate"`           // 전일 종가 대비 변화량
	ConvertedTradePrice  float64 `json:"converted_trade_price"` // 종가 환산 화폐 단위로 환산된 가격
}

/*
{
"market":"KRW-ETH",
"candle_date_time_utc":"2023-02-04T00:00:00",
"candle_date_time_kst":"2023-02-04T09:00:00",
"opening_price":2089000.00000000,
"high_price":2099000.00000000,
"low_price":2080000.00000000,
"trade_price":2087000.00000000,
"timestamp":1675483389080,
"candle_acc_trade_price":9743880531.14714000,
"candle_acc_trade_volume":4656.79844363,
"prev_closing_price":2091000.00000000,
"change_price":-4000.00000000,
"change_rate":-0.0019129603,
"converted_trade_price":null
}
*/

type GetWeekMonthCandleInputType struct {
	Market string
	To     string
	Count  int
}
type GetWeekMonthCandleInputParameterType struct {
	Market string
	To     string
	Count  string
}

type WeekMonthCandleInfo struct {
	Market               string  `json:"market"`
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`
	OpeningPrice         float64 `json:"opening_price"`
	HighPrice            float64 `json:"high_price"`
	LowPrice             float64 `json:"low_price"`
	TradePrice           float64 `json:"trade_price"`
	Timestamp            int     `json:"timestamp"`
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"`
	FirstDayOfPeriod     string  `json:"first_day_of_period"`
}

/*
{
"market":"KRW-ETH",
"candle_date_time_utc":"2023-01-30T00:00:00",
"candle_date_time_kst":"2023-01-30T09:00:00",
"opening_price":2054000.00000000,
"high_price":2138000.00000000,
"low_price":1938000.00000000,
"trade_price":2091000.00000000,
"timestamp":1675494991507,
"candle_acc_trade_price":329885745085.02079500,
"candle_acc_trade_volume":162009.99804735,
"first_day_of_period":"2023-01-30"
}
*/
