package Chart

import (
	"log"
	"upbit/Global"
)

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

// 추가적인 데이터 처리의 경우 차트 알고리즘 데이터에 넣는 기준은
// 기본 정보 수집 단위에 해당하는 주기가 된 실시간 틱 데이터인 경우 추가한다.
func GetFibonacciRetracementMACD(wallet ChartWallet, chartDataArr []Global.ChartDataForm) *FibonacciRetracement_MACD_type {
	// 계산 값이 담긴 객체의 포인터를 봇에다가 저장해야겠다.
	frmt := FibonacciRetracement_MACD_type{}
	frmt.wallet = wallet // 사용 할 금액 등 정보 기입
	frmt.wallet.lastBuyPrice = 0.0
	frmt.wallet.flag = true
	frmt.chartData = chartDataArr      // 차트 데이터 객체 내에 저장
	frmt.getFibonacciRetracementMACD() // 매수/매도 신호를 위한 기본 데이터 세팅

	return &frmt
}
func (f *FibonacciRetracement_MACD_type) FibonacciRetracementMACD(chartData Global.ChartDataForm) {
	f.checkChartDataAdd(chartData)
	f.runFibonacciRetracementMACD(chartData)
}

// 스톡캐스틱 지표는 변동성이 너무 크기 때문에, 추세 지표(MACD) 와 결합해서
// 사용해야하는 지표이다. 우선 알고리즘을 위해 기본 스톡캐스틱 알고리즘을 적용한다.
// (Kvalue와 Dvaule 가 교차하는 지점이 이벤트 발생 구간 -> 골든/데드 크로스)
// 추후 MACD 결합 시 더 다양한 조건이 붙을 예정
// 피보나치는 함수에 내장되어 MACD 가 적용 되어있는데
// 아래로 함수로 따로 빼서 피보나치 및 스톡캐스틱 이벤트에 적용 가능하도록
// 모듈화 필요 (타입, 함수 단위)
func GetStochastic(wallet ChartWallet, chartDataArr []Global.ChartDataForm) *Stochastic_type {
	st := Stochastic_type{}
	st.wallet = wallet // 사용 할 금액 등 정보 기입
	st.wallet.lastBuyPrice = 0.0
	st.wallet.flag = true
	st.chartData = chartDataArr
	// 아래 Kvalue 및 Dvalue 도 사용자에게서 받은 값으로 세팅 할 것인지 ?
	st.Krange = 5
	st.Drange = 3
	st.getStochastic()

	log.Printf("%+v", st)

	return &st
}
func (s *Stochastic_type) Stochastic(chartData Global.ChartDataForm) {
	// 스톡캐스틱은 데이터 추가를 어떻게 해야하는감 ?
	// -> 피보나치처럼 해당 주기가 되면, 스톡캐스틱 로직을 실행시켜주면 되나 ?
	// dValueArr 및 kValueArr 는 가장 뒷 쪽 인덱스가 최근 데이터고
	// chartData 는 가장 앞 인덱스가 최근 데이터이다.
	// 둘 다 역정렬이 필요 (get)
	log.Printf("%+v", s)
	s.checkChartDataAdd(chartData)
	s.runStochastic(chartData)

	// 여기 작업 끝나고 알고리즘 동작 확인 되면
	// 1.MACD 랑 합쳐서 알고리즘 이벤트 발생시키기
	// 2.B-bot botType.go, botFunc.go(getChartDataFromAbot) 부분 각 차트별 통합 방식 구현하기
}

func Bolingerband() {}
