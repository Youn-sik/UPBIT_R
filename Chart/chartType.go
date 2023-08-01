package Chart

type ChartWallet struct { // 하나의 차트 알고리즘에서 사용할 금액 정보
	TotalAmount  float64 // 투자 총 금액 (해당 봇의 해당 차트 지정 건)
	CurrencyType string  // 투자 금액 화폐 종류

	investAmount  float64 // 투자 중 금액
	currentAmount float64 // 미 투자 금액

	// 아래 두 개 데이터는 chart 패키지에서 관리할지, C-bot 에서 관리할지 ?
	lastBuyPrice float64 // 최근 투자 금액 (수익 데이터 집계에 사용)
	flag         bool    // 매수 flag, 매도 flag (true = 매수 Flag, false = 매도 Flag)

	timestamp   string // 투자 시작 일시
	uptimestamp string // 최근 투자 일시
}

type FibonacciRetracement_MACD_type struct {
	wallet ChartWallet

	maxPrice float64 // 종가 최대 값
	minPrice float64 // 종가 최소 값

	difference float64 // 종가 차잇 값

	zerothLevel float64 // 0% Level
	firstLevel  float64 // 23.6% Level
	secondLevel float64 // 38.2% Level
	thirdLevel  float64 // 50% Level
	fourthLevel float64 // 61.8% Level
	fifthLevel  float64 // 100% Level

	MACDInputValues []float64 // MACD input value
	macd            []float64 // MACD value
	signal          []float64 // MACD Signal value
	macdHistogram   []float64 // MACD Histogram value
}
