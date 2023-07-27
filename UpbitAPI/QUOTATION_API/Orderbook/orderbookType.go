package Orderbook

type Orderbook struct{}

type OrderbookInfo struct {
	Market         string          `json:"market"`
	Timestamp      int             `json:"timestamp"`
	TotalAskSize   float64         `json:"total_ask_size"`  // 호가 매도 총 잔량
	TotalBidSize   float64         `json:"total_bid_size"`  // 호가 매수 총 잔량
	OrderbookUnits []OrderbookUnit `json:"orderbook_units"` // 호가 (순서대로 1호가~15호가)
}
type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"` // 매도 호가
	BidPrice float64 `json:"bid_price"` // 매수 호가
	AskSize  float64 `json:"ask_size"`  // 매도 잔량
	BidSize  float64 `json:"bid_size"`  // 매수 잔량
}

/*
[
{
"market":"KRW-ETH",
"timestamp":1675517192989,
"total_ask_size":579.0202704499999,
"total_bid_size":283.51954158,
"orderbook_units":[
{
"ask_price":2129000.0,
"bid_price":2128000.0,
"ask_size":0.23785722,
"bid_size":0.95396471
},
{
"ask_price":2130000.0,
"bid_price":2127000.0,
"ask_size":188.30876992,
"bid_size":15.04351470
},
{
"ask_price":2131000.0,
"bid_price":2126000.0,
"ask_size":32.09900182,
"bid_size":20.13129092
},
{
"ask_price":2132000.0,
"bid_price":2125000.0,
"ask_size":44.40565110,
"bid_size":63.38632313
},
{
"ask_price":2133000.0,
"bid_price":2124000.0,
"ask_size":27.97415470,
"bid_size":19.12351980
},
{
"ask_price":2134000.0,
"bid_price":2123000.0,
"ask_size":12.50263607,
"bid_size":39.68566422
},
{
"ask_price":2135000.0,
"bid_price":2122000.0,
"ask_size":35.04876479,
"bid_size":6.04738170
},
{
"ask_price":2136000.0,
"bid_price":2121000.0,
"ask_size":6.68060746,
"bid_size":1.51077956
},
{
"ask_price":2137000.0,
"bid_price":2120000.0,
"ask_size":26.15427121,
"bid_size":37.00384700
},
{
"ask_price":2138000.0,
"bid_price":2119000.0,
"ask_size":88.87207859,
"bid_size":0.63424589
},
{
"ask_price":2139000.0,
"bid_price":2118000.0,
"ask_size":10.97668694,
"bid_size":0.18057758
},
{
"ask_price":2140000.0,
"bid_price":2117000.0,
"ask_size":67.91509323,
"bid_size":1.84863346
},
{
"ask_price":2141000.0,
"bid_price":2116000.0,
"ask_size":9.61225803,
"bid_size":5.58937095
},
{
"ask_price":2142000.0,
"bid_price":2115000.0,
"ask_size":18.70001813,
"bid_size":57.71656171
},
{
"ask_price":2143000.0,
"bid_price":2114000.0,
"ask_size":9.53242124,
"bid_size":14.66386625
}
]
}
]
*/
