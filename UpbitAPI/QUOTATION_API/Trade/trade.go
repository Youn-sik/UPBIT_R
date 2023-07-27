package Trade

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"upbit/Logger"
)

func (*Trade) GetTradeTick(getTradeTickInputType GetTradeTickInputType) ([]TradeTickInfo, string) {
	var tradeTickInfo []TradeTickInfo

	getTradeTickInputParameterType := setTradeTickInputToParameterType(getTradeTickInputType)
	url := "https://api.upbit.com/v1/trades/ticks?market=" + getTradeTickInputParameterType.Market +
		"&count=" + getTradeTickInputParameterType.Count + "&cursor=" + getTradeTickInputParameterType.Cursor
	if getTradeTickInputParameterType.To != "" {
		url += "&to=" + getTradeTickInputParameterType.To
	} else if getTradeTickInputParameterType.DaysAgo != "" {
		url += "&daysAgo=" + getTradeTickInputParameterType.DaysAgo
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &tradeTickInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return tradeTickInfo, "tradeTickInfo json Unmarshal 에 실패하였습니다."
	}

	return tradeTickInfo, ""
}
