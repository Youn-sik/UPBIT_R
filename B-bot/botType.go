package B_bot

import (
	"upbit/Chart"
	"upbit/Global"
)

type B_botType struct {
	// field ...
	Wallet Chart.ChartWallet

	// 차트 정보 포인터 (종류별로 처리 필요)
	chartPointer *Chart.FibonacciRetracement_MACD_type

	B_channel chan []Global.ChartDataForm
}
