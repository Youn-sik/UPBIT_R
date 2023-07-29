package Ticker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"upbit/Logger"
)

func (*Ticker) GetTicker(markets []string) ([]TickerInfo, string) {
	var tickerInfo []TickerInfo

	marketStr := setTickerMarketArrToStr(markets)

	url := "https://api.upbit.com/v1/ticker?markets=" + marketStr

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &tickerInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return tickerInfo, "tickerInfo json Unmarshal 을 실패하였습니다"
	}

	return tickerInfo, ""
}
