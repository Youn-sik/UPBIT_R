package Chart

import (
	"log"
	"time"
	"upbit/Global"

	"github.com/markcheno/go-talib"
)

func (f *FibonacciRetracement_MACD_type) getFibonacciRetracementMACD(chartDataArr []Global.ChartDataForm) {
	// wallet 상세 세팅
	ntime := time.Now().Format("2006-01-02T15:04:05")
	f.wallet.uptimestamp = ntime
	f.wallet.timestamp = ntime
	f.wallet.currentAmount = f.wallet.TotalAmount // 최초 전체 금액 = 현재 금액(투자하지 않은)

	// 1. 분석하는 기간의 주식 종가 최고가와 최저가 간의 차를 구함
	// 2. 차에서 각 레벨에 해당하는 퍼센티지를 곱하고 이를 최고가에서 뺌
	f.maxPrice = 0.0              // 최고가 비교 값
	f.minPrice = 99999999999999.0 // 최저가 비교 값

	for _, chart := range chartDataArr {
		op := chart.OpeningPrice
		// MACD를 위한 데이터를 초기화 해 준다.
		f.MACDInputValues = append(f.MACDInputValues, op)

		// 피보나치 값 구하기 위한 처리
		if f.maxPrice < op {
			f.maxPrice = op
		}
		if f.minPrice > op {
			f.minPrice = op
		}
	}
	f.difference = f.maxPrice - f.minPrice
	// 아래에서 현재 가격의 레벨을 구한다.
	// 각 레벨을 피보나치 값을 통해 지정한다.
	// 0% 는 maxPrice
	f.zerothLevel = f.maxPrice
	f.firstLevel = f.maxPrice - f.difference*0.236  // 23.6%
	f.secondLevel = f.maxPrice - f.difference*0.382 // 38.2%
	f.thirdLevel = f.maxPrice - f.difference*0.5    // 50%
	f.fourthLevel = f.maxPrice - f.difference*0.618 // 61.8%
	f.fifthLevel = f.minPrice
	// 100% 는 minPrice
	// 피보나치 되돌림 레벨 값 구하기 완료.
	/*
		log.Println("0%:", f.zerothLevel)
		log.Println("23.6%:", f.firstLevel)
		log.Println("38.2%:", f.secondLevel)
		log.Println("50%:", f.thirdLevel)
		log.Println("61.8%:", f.fourthLevel)
		log.Println("100%:", f.fifthLevel)
	*/

	// 아래는 MACD 값 구하기 (단기 이동평균선:12, 장기 이동평균선:26, 시그널:9 의 비율로 진행)
	// 기존 기준은 날짜이다.
	/*
		MACD 선 = 단기 지수 이동평균 - 장기 지수 이동평균
		신호선(Signal line) = MACD선의 지수 이동평균
		MACD Histogram(Oscillator) = MACD - 신호선 (더 빠른 반응 속도를 보여줌)
	*/
	periodShort := 12
	periodLong := 26
	periodSignal := 9

	// macd, signal, macdHistogram의 초반 값들이 0으로 나오는 이유는
	// 해당 틱에 대한 (틱의 기간) 계산 할 값이 마땅치 않아서 이다.

	f.macd, f.signal, f.macdHistogram = talib.Macd(f.MACDInputValues, periodShort, periodLong, periodSignal)
	// MACD, signal, MACDHistogram 값 구하기 완료
	/*
		log.Println(nowPrice)
		log.Println(MACDInputValues)
		log.Println(macd)
		log.Println(signal)
		log.Println(macdHistogram)
		log.Println(len(MACDInputValues), len(macd), len(signal), len(macdHistogram))
	*/
}

func (f *FibonacciRetracement_MACD_type) runFibonacciRetracementMACD(chartData Global.ChartDataForm) {
	f.wallet.flag = true        // true = 매수 Flag, false = 매도 Flag
	f.wallet.lastBuyPrice = 0.0 // 매도 시 매수 한 금액과 비교하기 위한 변수

	// f 에서 각종 정보를 가지고
	// 매개 변수의 chartData 를 현재값으로 flag 조정

	// 현재 금액을 기준으로 레벨을 얻는다
	nowPrice := chartData.OpeningPrice
	upperLevel, lowerLevel := f.getLevel(nowPrice)

	// 아래 시그널 파악 타이밍을 어떻게 잡을지 고민이다.
	// 기존 데이터의 가장 최근 값을 가지고 판단하여, 현재 가격으로 구매를 신청해야하는지 ? -> 우선 해당 방법으로 진행 (보다 쉬움)
	// 기존 데이터 + 현재 데이터 값을 가지고 판단하여, 현재 가격으로 구매를 신청해야하는지 ?
	lastMACDIndex := len(f.macd) - 1
	lastSignalIndex := len(f.signal) - 1
	if f.signal[lastSignalIndex] < f.macd[lastMACDIndex] && f.wallet.flag == true { // -> 여기서 bot C 에게 이벤트 전송해야한다.
		// 골든 시그널
		// MACD 라인이 Signal 라인을 넘게되면 buy signal 을 생성한다.
		// 이 방식은 한 개를 사면 산 것을 팔 때까지 다시 못 산다.
		f.wallet.lastBuyPrice = nowPrice

		/* -> 아래는 기록 남기는 부분으로 mongodb 에 저장이 필요하다.
		buyList = append(buyList, price)
		sellList = append(sellList, 0)
		*/
		f.wallet.flag = false
		f.wallet.currentAmount -= nowPrice
		f.wallet.investAmount += nowPrice

	} else if f.signal[lastSignalIndex] > f.macd[lastMACDIndex] && f.wallet.flag == false { // -> 여기서 bot C 에게 이벤트 전송해야한다.
		// } else if signal[idx] > macd[idx] && flag == false && price >= lastBuyPrice { // 해당 조건문은 이익이 나지 않으면 팔지 않는 조건문이다.
		// 데드 시그널

		/* -> 아래는 기록 남기는 부분으로 mongodb 에 저장이 필요하다.
		buyList = append(buyList, 0)
		sellList = append(sellList, price)
		*/
		f.wallet.flag = true
		f.wallet.currentAmount += nowPrice
		f.wallet.investAmount -= nowPrice

	} else if (f.wallet.lastBuyPrice >= upperLevel || f.wallet.lastBuyPrice <= lowerLevel) && f.wallet.flag == false { // -> 여기서 bot C 에게 이벤트 전송해야한다.
		// 데드 시그널 뿐만 아니라 피보나치 레벨에 도달한 경우 판매한다.
		// 해당 구매 가격의 윗 레벨인 경우 적당히 먹고 빠지는 것이며
		// 해당 구매 가격의 아래 레벨인 경우 더 잃기 전에 빠지는 거이다.

		/* -> 아래는 기록 남기는 부분으로 mongodb 에 저장이 필요하다.
		buyList = append(buyList, 0)
		sellList = append(sellList, price)
		*/
		f.wallet.flag = true
		f.wallet.currentAmount += nowPrice
		f.wallet.investAmount -= nowPrice
	}

	ntime := time.Now().Format("2006-01-02T15:04:05")
	f.wallet.uptimestamp = ntime
	log.Printf("wallet info: %+v", f.wallet)
}

// 이 함수는 현재 주식의 종가(오픈가)가 Fibonacci Retracement의 어느 레벨이 있는지를 확인한다.
// 현재 가격(오픈가, 종가)가 특정 레벨보다 높거나 같으면 그 레벨과 바로 아래의 하위 레벨을 불러온다.
// 해당 두 레벨을 통해 매도/매도 시점으로 비교가 가능하다.
func (f *FibonacciRetracement_MACD_type) getLevel(price float64) (float64, float64) {
	if price >= f.firstLevel {
		return f.zerothLevel, f.firstLevel
	} else if price >= f.secondLevel {
		return f.firstLevel, f.secondLevel
	} else if price >= f.thirdLevel {
		return f.secondLevel, f.thirdLevel
	} else if price >= f.fourthLevel {
		return f.thirdLevel, f.fourthLevel
	} else {
		return f.fourthLevel, f.fifthLevel
	}
}