package Global

import (
	"upbit/UpbitAPI/EXCHANGE_API/Account"
	"upbit/UpbitAPI/EXCHANGE_API/Authorization"
	"upbit/UpbitAPI/EXCHANGE_API/Order"
	"upbit/UpbitAPI/QUOTATION_API/Candle"
	"upbit/UpbitAPI/QUOTATION_API/Market"
	"upbit/UpbitAPI/QUOTATION_API/Orderbook"
	"upbit/UpbitAPI/QUOTATION_API/Ticker"
	"upbit/UpbitAPI/QUOTATION_API/Trade"
)

type UpbitClient struct {
	QUOTATION_API struct {
		Market    *Market.Market
		Candle    *Candle.Candle
		Orderbook *Orderbook.Orderbook
		Ticker    *Ticker.Ticker
		Trade     *Trade.Trade
	}
	EXCHANGE_API struct {
		Account       *Account.Account
		Authorization *Authorization.Authorization
		Order         *Order.Order
	}
}
