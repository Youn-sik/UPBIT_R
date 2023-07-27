package Market

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"upbit/Logger"
	"upbit/UpbitAPI/EXCHANGE_API/Authorization"
)

var auth Authorization.Authorization

func (*Market) GetMarketList() []MarketInfo {
	var marketInfo []MarketInfo

	url := "https://api.upbit.com/v1/market/all?isDetails=false"
	authorizationToken := auth.GetAuthoriztionToken("")

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authorizationToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &marketInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return marketInfo
	}

	return marketInfo
}
