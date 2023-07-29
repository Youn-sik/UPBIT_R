package A_bot

type A_botType struct {
	collectionInterval int
	marketName         string
	candleType         string // minute, day, week, month
	candleCount        int    // max 200
	candleUnit         int    // only minute candel -> 1,3,5,15,10,30,60,240

	// A-bot 에서 가공한 데이터
	chartData []ChartDataForm

	// 외부 객체 또는 다른 bot 에서 A 객체에게 데이터를 제공해야하면 아래 채널 활성화 하여 사용
	// A_channel chan TYPE
}

type ChartDataForm struct {
	Market       string
	OpeningPrice float64
	HighPrice    float64
	LowPrice     float64
	TradePrice   float64
	TimeKST      string
}
