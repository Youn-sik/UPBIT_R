package A_bot

import (
	"errors"
	"log"
	"upbit/Global"
	"upbit/Logger"
	"upbit/UpbitAPI/QUOTATION_API/Candle"
)

func (a *A_botType) run() {
	var candleInfo interface{}
	var errStr string

	switch a.candleType {
	case "minute":
		minCandleInput := getMinCandleInput(a.marketName, a.candleCount, a.candleUnit)
		candleInfo, errStr = Global.UpClient.QUOTATION_API.Candle.GetMinCandle(minCandleInput)
	case "day":
		dayCandleInput := getDayCandleInput(a.marketName, a.candleCount)
		candleInfo, errStr = Global.UpClient.QUOTATION_API.Candle.GetDayCandle(dayCandleInput)
	case "week":
		weekCandleInput := getWeekMonthCandleInput(a.marketName, a.candleCount)
		candleInfo, errStr = Global.UpClient.QUOTATION_API.Candle.GetWeekCandleInfo(weekCandleInput)
	case "month":
		monthCandleInput := getWeekMonthCandleInput(a.marketName, a.candleCount)
		candleInfo, errStr = Global.UpClient.QUOTATION_API.Candle.GetWeekCandleInfo(monthCandleInput)
	}

	if errStr != "" {
		Logger.PrintErrorLogLevel2(errors.New(errStr))
	}
	// B-bot 에게 candleInfo 제공
	log.Println(candleInfo)

}

func (a *A_botType) setDefaultConfig() {
	a.collectionInterval = 10
	a.marketName = "KRW-ETH"
	a.candleType = "minute"
	a.candleCount = 200
	a.candleUnit = 5
}

func (a *A_botType) setConfig(interval int, marketName, candleType string, candleCount, candleUnit int) {
	a.collectionInterval = interval
	a.marketName = marketName
	a.candleType = candleType
	a.candleCount = candleCount
	a.candleUnit = candleUnit
}

func getMinCandleInput(marketName string, candleCount, candleUnit int) Candle.GetMinCandleInputType {
	var minCandleInput Candle.GetMinCandleInputType
	minCandleInput.Market = marketName
	minCandleInput.Count = candleCount
	minCandleInput.MinUnit = candleUnit
	minCandleInput.To = "" // Latest Candle Data
	return minCandleInput
}

func getDayCandleInput(marketName string, candleCount int) Candle.GetDayCandleInputType {
	var dayCandleInput Candle.GetDayCandleInputType
	dayCandleInput.Market = marketName
	dayCandleInput.Count = candleCount
	dayCandleInput.To = ""                  // Latest Candle Data
	dayCandleInput.ConvertingPriceUnit = "" // 종가 환산 화폐 (마켓 코드를 원화 마켓이 아닌 곳에 지정한 경우 KRW로 사용)
	return dayCandleInput
}

func getWeekMonthCandleInput(marketName string, candleCount int) Candle.GetWeekMonthCandleInputType {
	var weekMonthCandleInput Candle.GetWeekMonthCandleInputType
	weekMonthCandleInput.Market = marketName
	weekMonthCandleInput.Count = candleCount
	weekMonthCandleInput.To = "" // Latest Candle Data
	return weekMonthCandleInput
}
