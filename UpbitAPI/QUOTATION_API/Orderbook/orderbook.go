package Orderbook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"upbit/Logger"
)

func (*Orderbook) GetOrderbook(markets []string) ([]OrderbookInfo, string) {
	var orderbookInfo []OrderbookInfo

	marketStr := setOrderbookMarketArrToStr(markets)

	url := "https://api.upbit.com/v1/orderbook?markets=" + marketStr

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &orderbookInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return orderbookInfo, "Orderbook json Unmarshal 에 실패하였습니다."
	}

	return orderbookInfo, ""
}
