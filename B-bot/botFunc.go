package B_bot

import (
	"log"
	"upbit/Chart"
	"upbit/Global"
)

func (b *B_botType) run() {
	// 버퍼 크기를 0으로 하면 수신처에서 받지 않으면 송신처에서도 대기한다. (블러킹 된다.)
	// non blocking 으로 진행하려면 버퍼를 지정해야한다.
	b.B_channel = make(chan []Global.ChartDataForm, 1)
	go b.getChartDataFromAbot()
}

func (b *B_botType) getChartDataFromAbot() {
	for chartData := range b.B_channel {
		log.Println("receiving data from B-bot")
		log.Printf("%+v", chartData)

		if len(chartData) != 1 { // 초기 데이터
			b.chartPointer = Chart.GetFibonacciRetracementMACD(b.Wallet, chartData)
			// b.chartPointer = Chart.GetStochastic(b.Wallet, chartData)
		} else { // 주기에 따른 틱 실시간 데이터
			b.chartPointer.FibonacciRetracementMACD(chartData[0])
			// b.chartPointer.Stochastic(chartData[0])
		}
	}
}
