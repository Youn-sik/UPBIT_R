package Ticker

type Ticker struct{}

type TickerInfo struct {
	Market             string  `json:"market"`
	TradeDate          string  `json:"trade_date"`      // 최근 거래 일자(UTC)
	TradeTime          string  `json:"trade_time"`      // 최근 거래 시각(UTC)
	TradeDateKST       string  `json:"trade_date_kst"`  // 최근 거래 일자(KST)
	TradeTimeKST       string  `json:"trade_time_kst"`  // 최근 거래 시각(KST)
	TradeTimestamp     int     `json:"trade_timestamp"` // 최근 거래 일시(UTC)
	OpeningPrice       int     `json:"opening_price"`
	HighPrice          int     `json:"high_price"`
	LowPrice           int     `json:"low_price"`
	TradePrice         int     `json:"trade_price"`           // 종가(현재가)
	PrevClosingPrice   float64 `json:"prev_closing_price"`    // 전일 종가
	Change             string  `json:"change"`                // EVEN:보합, RISE:상승, FALL:하락 (전일 종가 대비)
	ChangePrice        float64 `json:"change_price"`          // 변화액의 절대값 (전일 종가 대비)
	ChangeRate         float64 `json:"change_rate"`           // 변화율의 절대값 (전일 종가 대비)
	SignedChangePrice  float64 `json:"signed_change_price"`   // 부호가 있는 변화액 (전일 종가 대비)
	SignedChargeRate   float64 `json:"signed_change_rate"`    // 부호가 있는 변화액 (전일 종가 대비)
	TradeVolume        float64 `json:"trade_volume"`          // 가장 최근 거래량
	AccTradePrice      float64 `json:"acc_trade_price"`       // 누적 거래 대금(UTC 0시 기준)
	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`   // 24시간 누적 거래 대금
	AccTradeVolume     float64 `json:"acc_trade_volume"`      // 누적 거래량(UTC 0시 기준)
	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`  // 24시간 누적 거래량
	Highest52WeekPrice float64 `json:"highest_52_week_price"` // 52주 신고가
	Highest52WeekDate  string  `json:"highest_52_week_date"`  // 52주 신고가 달성일
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`  // 52주 신저가
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`   // 52주 신저가 달성일
	Timestamp          int     `json:"timestamp"`
}

/*
{
"market":"KRW-ETH",
"trade_date":"20230204",
"trade_time":"102335",
"trade_date_kst":"20230204",
"trade_time_kst":"192335",
"trade_timestamp":1675506215918,
"opening_price":2089000,
"high_price":2099000,
"low_price":2080000,
"trade_price":2095000,
"prev_closing_price":2091000.00000000,
"change":"RISE",
"change_price":4000.00000000,
"change_rate":0.0019129603,
"signed_change_price":4000.00000000,
"signed_change_rate":0.0019129603,
"trade_volume":2,
"acc_trade_price":15275113293.84005000,
"acc_trade_price_24h":48361389929.82538000,
"acc_trade_volume":7305.48117749,
"acc_trade_volume_24h":23228.93558072,
"highest_52_week_price":4348000.00000000,
"highest_52_week_date":"2022-04-03",
"lowest_52_week_price":1201500.00000000,
"lowest_52_week_date":"2022-06-18",
"timestamp":1675506216003
}
*/
