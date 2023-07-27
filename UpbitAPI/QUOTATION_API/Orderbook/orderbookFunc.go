package Orderbook

func setOrderbookMarketArrToStr(markets []string) string {
	var marketStr string

	for idx, val := range markets {
		if idx == 0 {
			marketStr += val
		} else {
			marketStr += "," + val
		}
	}

	return marketStr
}
