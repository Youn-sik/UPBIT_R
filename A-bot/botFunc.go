package A_bot

import (
	"errors"
	"log"
	"time"
	"upbit/Global"
	"upbit/Logger"
	"upbit/UpbitAPI/QUOTATION_API/Candle"
	"upbit/UpbitAPI/QUOTATION_API/Ticker"
)

func (a *A_botType) run() {
	var chartDataForm []ChartDataForm
	var errStr string

	switch a.candleType {
	case "minute":
		// 분 단위의 캔들을 가져오기 위해 input 값을 정의한다.
		minCandleInput := getMinCandleInput(a.marketName, a.candleCount, a.candleUnit)
		// 정의한 input 을 사용하여 분 단위 캔들을 가져온다.
		candleInfo, errstr := Global.UpClient.QUOTATION_API.Candle.GetMinCandle(minCandleInput)
		// 분 단위 캔들 정보를 A-bot 의 데이터 폼에 맞춘다.
		chartDataForm = convertCandleTypeToAbotDataForm(candleInfo)
		// switch case 문 밖에서 에러 처리를 위해 아래 코드를 사용한다.
		errStr = errstr
	case "day":
		dayCandleInput := getDayCandleInput(a.marketName, a.candleCount)
		candleInfo, errstr := Global.UpClient.QUOTATION_API.Candle.GetDayCandle(dayCandleInput)
		chartDataForm = convertCandleTypeToAbotDataForm(candleInfo)
		errStr = errstr
	case "week":
		weekCandleInput := getWeekMonthCandleInput(a.marketName, a.candleCount)
		candleInfo, errstr := Global.UpClient.QUOTATION_API.Candle.GetWeekCandleInfo(weekCandleInput)
		chartDataForm = convertCandleTypeToAbotDataForm(candleInfo)
		errStr = errstr
	case "month":
		monthCandleInput := getWeekMonthCandleInput(a.marketName, a.candleCount)
		candleInfo, errstr := Global.UpClient.QUOTATION_API.Candle.GetWeekCandleInfo(monthCandleInput)
		chartDataForm = convertCandleTypeToAbotDataForm(candleInfo)
		errStr = errstr
	}

	if errStr != "" {
		Logger.PrintErrorLogLevel2(errors.New(errStr))
	}

	// B-bot 에게 candleInfo 제공을 위해 객체 필드에 저장
	// log.Printf("%+v", botDataForm)
	a.addChartData(chartDataForm)
	// log.Printf("%+v", a.chartData)

	// B-bot 에게 제공할 데이터 형식으로 변환을 완료 하였으니,
	// 아래에서 B-bot 에게 해당 데이터를 넘겨준다.
	// B-bot 에게 데이터 넘겨주는 코드

	// ---
	// candleInfo 를 가져오고 난 후 지정한 주기마다 interval 을 돌면서 현재 호가 정보(중 현재가)를 불러온다.
	// 구매할 때, 호가창(order book)을 가져와서 ticker 에서 마지노선의 차이를 잡아서 주문하도록
	// interval => a.interval ticker 요청
	// 여기서 인터발 함수 실행

	// initalize channel
	tickChannel := make(chan Ticker.TickerInfo)

	go a.runIntervalTickerRequest(tickChannel)
	// 아래 함수에서 channel 로 부터 현재가 정보를 받는다.
	// 위에서 작업 한 B-bot 에게 데이터 전달을 위한 형식 변경을 아래 함수에서도 진행한다.
	// 그 후 B-bot 에게 tickInfo 제공
	a.getIntervalTickerResponse(tickChannel)
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

func (a *A_botType) runIntervalTickerRequest(ch chan Ticker.TickerInfo) {

	tickerInterval := time.NewTicker(time.Duration(a.collectionInterval) * time.Second)
	defer tickerInterval.Stop()

	var tickerInfo []Ticker.TickerInfo
	for range tickerInterval.C {
		var errStr string
		tickerInfo, errStr = Global.UpClient.QUOTATION_API.Ticker.GetTicker([]string{a.marketName})
		if errStr != "" {
			Logger.PrintErrorLogLevel2(errors.New(errStr))
		}
		// 여기서 a 객체에 정의된 채널로 데이터 전송
		ch <- tickerInfo[0]
	}

	// 채널을 정리(종료)하는 기능이 필요하다면
	// close(ch) 함수를 사용하면 된다.
}

func (a *A_botType) getIntervalTickerResponse(ch chan Ticker.TickerInfo) {
	for tick := range ch {
		// log.Printf("%+v", tick)
		chartDataForm := convertTickTypeToAbotDataForm(tick)
		log.Printf("%+v", chartDataForm)

		// 아래 이 함수를 쓰면 기존에 실행 시 저장되는
		// 지정한 단위의 캔들 정보에 현재 틱 정보를 넣는건데,
		// 그렇게 되면 전체 데이터를 또 한 번 B-bot 에게 전달해야한다.
		// 그 보다는 추가된 내용만 따로 전송하는게 나을 것 같다.
		// a.addChartData()
		// ---
		// B-bot 에게 데이터 넘겨주는 코드

	}
}

func convertCandleTypeToAbotDataForm[T Candle.CandleTypeGeneric](candle []T) []ChartDataForm {
	var chartDataForm []ChartDataForm

	// 위 객체의 timestampKST 변수에 string 형 한국 시간이라서
	// candle 패키지에서 받아온 timestamp 를 변환해 주기 위해 아래 함수를 사용해야한다.
	// Global.ConverTimestampTo
	for _, val := range candle {
		var cdf ChartDataForm
		// log.Printf("%+v", val)
		market, openingPrice, highPrice, lowPrice, tradePrice, timestamp := val.GetCandleInfo()
		timeKST := Global.ConvertTimestampToKST(timestamp)

		cdf.Market = market
		cdf.OpeningPrice = openingPrice
		cdf.HighPrice = highPrice
		cdf.LowPrice = lowPrice
		cdf.TradePrice = tradePrice
		cdf.TimeKST = timeKST

		chartDataForm = append(chartDataForm, cdf)
	}

	return chartDataForm
}

func convertTickTypeToAbotDataForm(tick Ticker.TickerInfo) []ChartDataForm {
	var cdf ChartDataForm

	timeKST := Global.ConvertTimestampToKST(tick.Timestamp)
	cdf.Market = tick.Market
	cdf.OpeningPrice = tick.OpeningPrice
	cdf.HighPrice = tick.HighPrice
	cdf.LowPrice = tick.LowPrice
	cdf.TradePrice = tick.TradePrice
	cdf.TimeKST = timeKST

	return []ChartDataForm{cdf}
}

func (a *A_botType) addChartData(chartDataForm []ChartDataForm) {
	a.chartData = append(chartDataForm)
}
