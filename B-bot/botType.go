package B_bot

type B_botType struct {
	// field ...

	// channel (*사용 시 make 필수*)
	B_channel chan []ChartDataForm
}

type ChartDataForm struct {
	Market       string
	OpeningPrice float64
	HighPrice    float64
	LowPrice     float64
	TradePrice   float64
	TimeKST      string
}
