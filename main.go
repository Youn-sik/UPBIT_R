package main

import (
	A_bot "upbit/A-bot"
)

func main() {
	// 실행

	// aBot1 := A_bot.NewBot()
	// aBot1.SetDefaultConfig()
	// aBot1.SetConfig(10, "KRW-ETH", "minute", 200, 5)
	// aBot1.SetConfig(10, "KRW-BTC", "day", 200, 0)
	// aBot1.Run()

	aBot2 := A_bot.NewBot()
	// aBot2.SetDefaultConfig()
	// aBot2.SetConfig(10, "KRW-ETH", "minute", 200, 5)
	// aBot2.SetConfig(10, "KRW-BTC", "day", 200, 0)

	// aBot2.SetConfig(10, "KRW-BTC", "week", 200, 0)
	aBot2.SetConfig(10, "KRW-BTC", "month", 200, 0)

	aBot2.Run()
}
