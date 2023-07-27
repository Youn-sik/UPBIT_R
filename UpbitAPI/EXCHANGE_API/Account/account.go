package Account

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"upbit/Logger"

	"upbit/UpbitAPI/EXCHANGE_API/Authorization"
)

var auth Authorization.Authorization

func (*Account) GetAccounts() []AccountInfo {
	var accountInfo []AccountInfo
	var authorizationToken = auth.GetAuthoriztionToken("")

	url := "https://api.upbit.com/v1/accounts"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authorizationToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &accountInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return accountInfo
	}

	return accountInfo
}
