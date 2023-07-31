package B_bot

import "upbit/Global"

type B_botType struct {
	// field ...

	// channel (*사용 시 make 필수*)
	B_channel chan []Global.ChartDataForm
}
