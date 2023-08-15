package main

import (
	A_bot "upbit/A-bot"
)

func main() {
	// 실행

	testBot := A_bot.NewBot()
	testBot.SetDefaultConfig()
	go testBot.Run()
	// testBot.Run()

	testBot1 := A_bot.NewBot()
	testBot1.SetConfig(10, "KRW-BTC", "minute", 200, 5, "KRW", 100000000)
	go testBot1.Run()
	// testBot1.Run()

	testBot2 := A_bot.NewBot()
	testBot2.SetConfig(10, "KRW-ETH", "minute", 200, 5, "KRW", 100000000)
	testBot2.Run()

	// 모든 run method 는 go 고루틴으로 실행 시키고
	// 클라이언트와 인터페이스가 되는 tcp 서버를 아래에서 실행 시켜두면 된다.
	// tcp 서버의 func 에서도 봇이 생성되어 실행 될 때, 고루틴으로 Run() 실행.
}

// B-bot 이 돌면서 알아야하는 데이터는 이전 데이터(위에서 지정한 주기에 따른 데이터 - 분,일,주,월) + 지정한 주기마다의 실시간 데이터
// 하지만 그러면 데이터들의 시간 차이 일관성이 깨진다. 그걸 맞춰줘야하긴 한다.
// 해당 작업에서 고려하게되면 단타와 장투를 같이 걸어두는 봇을 동시에 운영할 수 있게 된다.
// 우선 지정한 단위의 데이터를 가져오고 + 지정한 주기바다 받아오는 데이터와 결합해서 알고리즘에 적용하자.
// 그렇다면 A-bot 은 우선 지정한 단위의 데이터를 가져온 후에 interval 을 돌면서 현재 호가 정보를 주기에 맞춰 계속 가져와야한다.

// 지정된 주기에 따른 데이터(candleInfo)를 가져오고 난 후 지정한 주기마다 interval 을 돌면서 현재 호가 정보를 불러온다.
// 지정된 주기마다 호가 창 (판매 또는 구입을 위해 사용자들이 올려둔 금액의 리스트) 를 가져올지
// ticker = 시세 현재가(현재 종가) 를 가져올지 -> ticker 을 가져오고 그거로 계산하자
// 구매할 때, 호가창을 가져와서 ticker 에서 마지노선의 차이를 잡아서 주문하도록
