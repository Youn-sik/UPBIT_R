package Order

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"upbit/Logger"
	"upbit/UpbitAPI/EXCHANGE_API/Authorization"
)

var auth Authorization.Authorization

func (*Order) GetOrderChance(market string) OrderChanceInfo {
	var orderChanceInfo OrderChanceInfo

	queryString := "market=" + market
	authorizationToken := auth.GetAuthoriztionToken(queryString)
	url := "https://api.upbit.com/v1/orders/chance?" + queryString

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authorizationToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &orderChanceInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return orderChanceInfo
	}

	return orderChanceInfo
}

func (*Order) GetOrder(uuid, identifier string) OrderDetailInfo {
	var orderDetailInfo OrderDetailInfo

	queryString := setGetOrderQueryString(uuid, identifier)
	authorizationToken := auth.GetAuthoriztionToken(queryString)
	url := "https://api.upbit.com/v1/order" + queryString

	log.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authorizationToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// 결과값 json unmarshal 작업 필요
	log.Println(string(body))

	return orderDetailInfo
}

func (*Order) GetOrderList(
	market string,
	uuids, identifiers []string,
	state string,
	states []string,
	page, limit int,
	orderBy string,
) OrderInfo {
	var orderInfo OrderInfo

	queryString := setGetOrderListQueryString(market, uuids, identifiers, state, states, page, limit, orderBy)
	authorizationToken := auth.GetAuthoriztionToken(queryString)
	url := "https://api.upbit.com/v1/orders" + queryString

	log.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authorizationToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// 결과값 json unmarshal 작업 필요
	log.Println(string(body))

	return orderInfo
}

func (*Order) OrderCancel(uuid, identifier string) OrderCancelInfo {
	var orderCancelInfo OrderCancelInfo

	queryString := setOrderCancelQueryString(uuid, identifier)
	authorizationToken := auth.GetAuthoriztionToken(queryString)
	url := "https://api.upbit.com/v1/order" + queryString

	log.Println(url)

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authorizationToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	log.Println(string(body))

	return orderCancelInfo
}

func (*Order) OrderNew(market, side, volume, price, ordType, identifier string) OrderInfo {
	var orderInfo OrderInfo

	url := "https://api.upbit.com/v1/orders"

	payload := strings.NewReader(
		"{\"market\":\"" + market +
			"\",\"side\":\"" + side +
			"\",\"volume\":\"" + volume +
			"\",\"identifier\":\"" + identifier +
			"\",\"ord_type\":\"" + ordType +
			"\",\"price\":\"" + price + "\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	log.Println(string(body))

	return orderInfo
}
