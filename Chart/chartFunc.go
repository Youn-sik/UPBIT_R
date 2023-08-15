package Chart

import (
	"log"
	"sort"
	"time"
	"upbit/Global"
	"upbit/Logger"

	"github.com/markcheno/go-talib"
)

func (f *FibonacciRetracement_MACD_type) getFibonacciRetracementMACD() {
	// wallet 상세 세팅
	ntime := time.Now().Format("2006-01-02T15:04:05")
	f.wallet.uptimestamp = ntime
	f.wallet.timestamp = ntime
	f.wallet.currentAmount = f.wallet.TotalAmount // 최초 전체 금액 = 현재 금액(투자하지 않은)

	// 1. 분석하는 기간의 주식 종가 최고가와 최저가 간의 차를 구함
	// 2. 차에서 각 레벨에 해당하는 퍼센티지를 곱하고 이를 최고가에서 뺌
	f.maxPrice = 0.0              // 최고가 비교 값
	f.minPrice = 99999999999999.0 // 최저가 비교 값

	for _, chart := range f.chartData {
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
	// f 에서 각종 정보를 가지고
	// 매개 변수의 chartData 를 현재값으로 flag 조정

	// 현재 금액을 기준으로 레벨을 얻는다
	nowPrice := chartData.OpeningPrice
	upperLevel, lowerLevel := f.getLevel(nowPrice)

	// 아래 시그널 파악 타이밍을 어떻게 잡을지 고민이다.
	// 0번 인덱스가 가장 최근의 데이터
	if f.signal[0] < f.macd[0] && f.wallet.flag == true { // -> 여기서 bot C 에게 이벤트 전송해야한다.
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
		f.wallet.TotalAmount -= nowPrice
		f.wallet.investAmount += nowPrice

	} else if f.signal[0] > f.macd[0] && f.wallet.flag == false { // -> 여기서 bot C 에게 이벤트 전송해야한다.
		// } else if signal[idx] > macd[idx] && flag == false && price >= lastBuyPrice { // 해당 조건문은 이익이 나지 않으면 팔지 않는 조건문이다.
		// 데드 시그널

		/* -> 아래는 기록 남기는 부분으로 mongodb 에 저장이 필요하다.
		buyList = append(buyList, 0)
		sellList = append(sellList, price)
		*/
		f.wallet.flag = true
		f.wallet.currentAmount += f.wallet.lastBuyPrice
		f.wallet.TotalAmount += f.wallet.lastBuyPrice
		f.wallet.investAmount -= f.wallet.lastBuyPrice
		f.wallet.lastBuyPrice = 0.0

	} else if (f.wallet.lastBuyPrice >= upperLevel || f.wallet.lastBuyPrice <= lowerLevel) && f.wallet.flag == false { // -> 여기서 bot C 에게 이벤트 전송해야한다.
		// 데드 시그널 뿐만 아니라 피보나치 레벨에 도달한 경우 판매한다.
		// 해당 구매 가격의 윗 레벨인 경우 적당히 먹고 빠지는 것이며
		// 해당 구매 가격의 아래 레벨인 경우 더 잃기 전에 빠지는 거이다.

		/* -> 아래는 기록 남기는 부분으로 mongodb 에 저장이 필요하다.
		buyList = append(buyList, 0)
		sellList = append(sellList, price)
		*/
		f.wallet.flag = true
		f.wallet.currentAmount += f.wallet.lastBuyPrice
		f.wallet.TotalAmount += f.wallet.lastBuyPrice
		f.wallet.investAmount -= f.wallet.lastBuyPrice
		f.wallet.lastBuyPrice = 0.0
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

// 기본 정보 수집 데이터의 period 가 요청한 시기의 시간 (인덱스0) 이후 period 마다로 데이터가 와서 우짜지
func (f *FibonacciRetracement_MACD_type) checkChartDataAdd(chartData Global.ChartDataForm) {
	lastTimestamp1 := f.chartData[0].TimeKST // 마지막 정보의 수집 시간 정보를 가져온다. (하나 더 전. 왜냐면 제일 마지막 건 요청한 시간으로 들어온다)
	lastTimestamp2 := f.chartData[1].TimeKST // 마지막의 이전 정보의 수집 시간 정보를 가져온다. (하나 더 전. 왜냐면 제일 마지막 건 요청한 시간으로 들어온다)

	lt1, err := time.Parse("2006-01-02T15:04:05", lastTimestamp1)   // 기존 수집 데이터의 timestamp 를 time 형태로 변환
	lt2, err := time.Parse("2006-01-02T15:04:05", lastTimestamp2)   // 기존 수집 데이터의 timestamp 를 time 형태로 변환
	nt, err := time.Parse("2006-01-02T15:04:05", chartData.TimeKST) // 틱 정보의 timestamp 를 time 형태로 변환
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return
	}

	existingPeriod := lt1.Sub(lt2) // 기존 수집 데이터의 주기를 구한다.
	nowPeriod := nt.Sub(lt1)
	// if nowPeriod

	log.Println(existingPeriod)
	log.Println(nowPeriod)
	// 하지만 아래 조건일 경우, 봇을 오래 틀어둠에 따라 기존 수집 데이터의 시간 단위보다 시간 단위가 점점 커지게 된다.
	// 근본적인 해결 방안은 A-bot 에게로 부터 받아온 데이터를 사용(사용자가 지정한 최초 데이터의 시간 단위)

	if existingPeriod <= nowPeriod { // 기본 수집 데이터의 시간 단위보다 같거나 이후의 데이터인 경우 알고리즘 데이터에 반영한다.
		log.Println("yy")
		// f.chartData = append(f.chartData, chartData) // append 가 아니라 prepend 가 되어야 한다.
		f.chartData = append([]Global.ChartDataForm{chartData}, f.chartData...)
		f.getFibonacciRetracementMACD()
	}
}

func (s *Stochastic_type) getStochastic() {
	// wallet 상세 세팅
	ntime := time.Now().Format("2006-01-02T15:04:05")
	s.wallet.uptimestamp = ntime
	s.wallet.timestamp = ntime
	s.wallet.currentAmount = s.wallet.TotalAmount // 최초 전체 금액 = 현재 금액(투자하지 않은)

	// 객체에 chartData, wallet, Kvalue/Dvalue 데이터 존재
	s.compareArr = make([]float64, s.Krange)
	// 근데 아래역 정렬이 최초 한 번만 되어야하나 ? 이후 추가되는 데이터는 어떻게 ?
	// 시간 기준 최근게 마지막으로 정렬
	sort.Slice(s.chartData, func(i, j int) bool {
		t1, err := time.Parse("2006-01-02T15:04:05", s.chartData[i].TimeKST)
		if err != nil {
			log.Println(err)
		}
		t2, _ := time.Parse("2006-01-02T15:04:05", s.chartData[j].TimeKST)
		return t1.Before(t2)
	})

	for idx, val := range s.chartData {
		var nowPrice float64 = val.OpeningPrice
		var kValue, minPrice, maxPrice float64

		// Krange 내에서 최고값, 최저값을 가져온다.
		s.compareArr, minPrice, maxPrice = setCompareArr(s.compareArr, nowPrice, idx, s.Krange)
		s.openingPriceHighArr = append(s.openingPriceHighArr, maxPrice)
		s.openingPriceLowArr = append(s.openingPriceLowArr, minPrice)

		// Kvalue 를 구하려면 해당 날로부터 이전 날의 Krange 만큼의 데이터가 있어야한다.
		if idx >= s.Krange {
			// %K = (현재 가격의 폭) / (파동의 전체 폭) * 100(%)
			kValue = (nowPrice - minPrice) / (maxPrice - minPrice) * 100
			s.kValueArr = append(s.kValueArr, kValue)
		} else {
			s.kValueArr = append(s.kValueArr, 0)
		}
	}
	// Dvalue 는 KvalueArr 의 Drange 만큼의 이평선 값이다.
	s.dValueArr = movingAverage(s.kValueArr, s.Drange)

	// 시간 기준 최근게 앞으로 정렬
	sort.Slice(s.chartData, func(i, j int) bool {
		t2, err := time.Parse("2006-01-02T15:04:05", s.chartData[i].TimeKST)
		if err != nil {
			log.Println(err)
		}
		t1, _ := time.Parse("2006-01-02T15:04:05", s.chartData[j].TimeKST)
		return t1.Before(t2)
	})
}

func setCompareArr(compareArr []float64, nowPrice float64, idx, Krange int) ([]float64, float64, float64) {
	var maxPrice, minPrice float64 = 0, 999999999999 // 9천억

	if idx == 0 {
		for i := 0; i < len(compareArr); i++ {
			compareArr[i] = nowPrice
		}
	} else {
		compareArr[idx%Krange] = nowPrice
	}

	for _, val := range compareArr {
		if val <= minPrice {
			minPrice = val
		}
		if val >= maxPrice {
			maxPrice = val
		}
	}

	return compareArr, minPrice, maxPrice
}

func movingAverage(data []float64, windowSize int) []float64 {
	if windowSize > len(data) {
		windowSize = len(data)
	}

	var result []float64
	for idx, _ := range data {
		if idx < windowSize-1 {
			result = append(result, 0)
		} else {
			var mv float64
			for ws := windowSize - 1; ws > 0; ws-- { // window size == 5 인 경우 4,3,2,1
				mv += data[idx-ws]
			}
			rst := mv / float64(windowSize)
			result = append(result, rst)
		}
	}

	return result
}

func (s *Stochastic_type) checkChartDataAdd(chartData Global.ChartDataForm) {
	lastTimestamp1 := s.chartData[0].TimeKST // 마지막 정보의 수집 시간 정보를 가져온다. (하나 더 전. 왜냐면 제일 마지막 건 요청한 시간으로 들어온다)
	lastTimestamp2 := s.chartData[1].TimeKST // 마지막의 이전 정보의 수집 시간 정보를 가져온다. (하나 더 전. 왜냐면 제일 마지막 건 요청한 시간으로 들어온다)

	lt1, err := time.Parse("2006-01-02T15:04:05", lastTimestamp1)   // 기존 수집 데이터의 timestamp 를 time 형태로 변환
	lt2, err := time.Parse("2006-01-02T15:04:05", lastTimestamp2)   // 기존 수집 데이터의 timestamp 를 time 형태로 변환
	nt, err := time.Parse("2006-01-02T15:04:05", chartData.TimeKST) // 틱 정보의 timestamp 를 time 형태로 변환
	if err != nil {
		Logger.PrintErrorLogLevel2(err)
		return
	}

	existingPeriod := lt1.Sub(lt2) // 기존 수집 데이터의 주기를 구한다.
	nowPeriod := nt.Sub(lt1)
	// if nowPeriod

	log.Println(existingPeriod)
	log.Println(nowPeriod)
	// 하지만 아래 조건일 경우, 봇을 오래 틀어둠에 따라 기존 수집 데이터의 시간 단위보다 시간 단위가 점점 커지게 된다.
	// 근본적인 해결 방안은 A-bot 에게로 부터 받아온 데이터를 사용(사용자가 지정한 최초 데이터의 시간 단위)

	if existingPeriod <= nowPeriod { // 기본 수집 데이터의 시간 단위보다 같거나 이후의 데이터인 경우 알고리즘 데이터에 반영한다.
		log.Println("yy")
		// f.chartData = append(f.chartData, chartData) // append 가 아니라 prepend 가 되어야 한다.
		s.chartData = append([]Global.ChartDataForm{chartData}, s.chartData...)
		s.getStochastic()
	}
}

func (s *Stochastic_type) runStochastic(chartData Global.ChartDataForm) {
	// s 에서 각종 정보를 가지고
	// 매개 변수의 chartData 를 현재값으로 flag 조정

	nowPrice := chartData.OpeningPrice
	// ??? 만약 골든 크로스가 발생했다. 근데 그 정보는 차트 판단 주기가 돌아오기 이전까지
	// 계속 골든 크로스로 인식할텐데 ? 구매 플래그 때문에 상관 없다 ?
	//kValue, dValue 는 가장 마지막 인덱스가 최근 데이터이다.
	// 이 두개 데이터가 교차되는 지점을 알아야한다.
	lastKvalue := s.kValueArr[len(s.kValueArr)-1]       // kValue 가장 마지막 값
	lastKvalueBefore := s.kValueArr[len(s.kValueArr)-2] // kValue 가장 마지막의 이전 값
	lastDvalue := s.dValueArr[len(s.dValueArr)-1]       // dValue 가장 마지막 값
	lastDvalueBefore := s.dValueArr[len(s.dValueArr)-2] // dValue 가장 마지막의 이전 값

	// 추후에 MACD 추세 그래프까지 적용 필요
	if lastKvalueBefore < lastDvalueBefore && lastKvalue > lastDvalue && s.wallet.flag == true && lastKvalue > 20 { // -> 여기서 bot C 에게 이벤트 전송해야한다.
		// 골든 시그널 (20% 초과에서만)
		// 이 방식은 한 개를 사면 산 것을 팔 때까지 다시 못 산다.
		s.wallet.lastBuyPrice = nowPrice

		/* -> 아래는 기록 남기는 부분으로 mongodb 에 저장이 필요하다.
		buyList = append(buyList, price)
		sellList = append(sellList, 0)
		*/
		s.wallet.flag = false
		s.wallet.currentAmount -= nowPrice
		s.wallet.TotalAmount -= nowPrice
		s.wallet.investAmount += nowPrice

	} else if lastKvalueBefore > lastDvalueBefore && lastKvalue < lastDvalue && s.wallet.flag == false && lastKvalue < 80 { // -> 여기서 bot C 에게 이벤트 전송해야한다.
		// 데드 시그널 (80% 미만에서만)

		/* -> 아래는 기록 남기는 부분으로 mongodb 에 저장이 필요하다.
		buyList = append(buyList, 0)
		sellList = append(sellList, price)
		*/
		s.wallet.flag = true
		s.wallet.currentAmount += s.wallet.lastBuyPrice
		s.wallet.TotalAmount += s.wallet.lastBuyPrice
		s.wallet.investAmount -= s.wallet.lastBuyPrice
		s.wallet.lastBuyPrice = 0.0
	}

	ntime := time.Now().Format("2006-01-02T15:04:05")
	s.wallet.uptimestamp = ntime
	log.Printf("wallet info: %+v", s.wallet)
}
