package Chart

import "upbit/Global"

/*
A-bot의 candleType 에 따라 기초 데이터 수집이 다르게 온다.
chart 패키지에서는 B-bot의 ChartData 형식만을 가지고 계산하면 된다.
만약 시간에 대한 계산이 필요한 경우 확인이 필요하다.

1. 최초로 받아오는 데이터를 먼저 처리
2. 이후로 추가적으로 받아오는 데이터 처리
*/

// 아래 함수 B-bot 에서 사용하도록 작업 필요
// 다만 첫 번째 데이터 수집의 경우 GetFibonacciRetracementMACD 함수로,
// 이후 추가적인 데이터 수집의 경우 fibonacciRetracementMACD 함수로 요청 보내도록 해야한다.
func GetFibonacciRetracementMACD(wallet ChartWallet, chartDataArr []Global.ChartDataForm) *fibonacciRetracement_MACD_type {
	// 계산 값이 담긴 객체의 포인터를 봇에다가 저장해야겠다.
	frmt := fibonacciRetracement_MACD_type{}
	frmt.wallet = wallet                           // 사용 할 금액 등 정보 기입
	frmt.getFibonacciRetracementMACD(chartDataArr) // 매수/매도 신호를 위한 기본 데이터 세팅

	return &frmt
}
func (f *fibonacciRetracement_MACD_type) fibonacciRetracementMACD(chartData Global.ChartDataForm) {
	f.runFibonacciRetracementMACD(chartData)
}

func Stochastic()

func Bolingerband()
