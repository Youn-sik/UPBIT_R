package Candle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"upbit/Logger"
)

func (*Candle) GetMinCandle(getMinCandleInputType GetMinCandleInputType) ([]MinCandleInfo, string) {
	var minCandleInfo []MinCandleInfo

	// if getMinCandleInputCheck(getMinCandleInputType) {
	// 	return minCandleInfo, "올바르지 않은 입력 값입니다."
	// }

	getMinCandleInputParameterType := setMinCandleInputToParameterType(getMinCandleInputType)
	url := "https://api.upbit.com/v1/candles/minutes/" + getMinCandleInputParameterType.MinUnit +
		"?market=" + getMinCandleInputParameterType.Market + "&count=" + getMinCandleInputParameterType.Count
	if getMinCandleInputParameterType.To != "" {
		url += "&to=" + getMinCandleInputParameterType.To
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &minCandleInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return minCandleInfo, "minCandleInfo JSON Unmarshal 실패"
	}

	return minCandleInfo, ""
}

func (*Candle) GetDayCandle(getDayCandleInputType GetDayCandleInputType) ([]DayCandleInfo, string) {
	var dayCandleInfo []DayCandleInfo

	getDayCandleInputParameterType := setDayCandleInputToParameterType(getDayCandleInputType)
	url := "https://api.upbit.com/v1/candles/days?market=" + getDayCandleInputParameterType.Market +
		"&count=" + getDayCandleInputParameterType.Count
	if getDayCandleInputParameterType.To != "" {
		url += "&to=" + getDayCandleInputParameterType.To
	} else if getDayCandleInputParameterType.ConvertingPriceUnit != "" {
		url += "&convertingPriceUnit=" + getDayCandleInputParameterType.ConvertingPriceUnit
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &dayCandleInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return dayCandleInfo, "dayCandleInfo JSON Unmarshal 실패"
	}

	return dayCandleInfo, ""
}

func (*Candle) GetWeekCandleInfo(getWeekCandleInputType GetWeekMonthCandleInputType) ([]WeekMonthCandleInfo, string) {
	var weekCandleInfo []WeekMonthCandleInfo

	getWeekCandleInputParameterType := setWeekMonthCandleInputToParameterType(getWeekCandleInputType)
	url := "https://api.upbit.com/v1/candles/weeks?market=" + getWeekCandleInputParameterType.Market +
		"&count=" + getWeekCandleInputParameterType.Count
	if getWeekCandleInputParameterType.To != "" {
		url += "&to=" + getWeekCandleInputParameterType.To
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &weekCandleInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return weekCandleInfo, "weekCandleInfo JSON Unmarshal 실패"
	}

	return weekCandleInfo, ""
}

func (*Candle) GetMonthCandleInfo(getMonthCandleInputType GetWeekMonthCandleInputType) ([]WeekMonthCandleInfo, string) {
	var monthCandleInfo []WeekMonthCandleInfo

	getMonthCandleInputParameterType := setWeekMonthCandleInputToParameterType(getMonthCandleInputType)
	url := "https://api.upbit.com/v1/candles/months?market=" + getMonthCandleInputParameterType.Market +
		"&count=" + getMonthCandleInputParameterType.Count
	if getMonthCandleInputParameterType.To != "" {
		url += "&to=" + getMonthCandleInputParameterType.To
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &monthCandleInfo)
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return monthCandleInfo, "monthCandleInfo JSON Unmarshal 실패"
	}

	return monthCandleInfo, ""
}
