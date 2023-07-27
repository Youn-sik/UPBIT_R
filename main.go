package main

import (
	A_bot "upbit/A-bot"
)

func main() {
	// 실행
	aBot1 := A_bot.NewBot()
	aBot1.SetDefaultConfig()
	// 또는 aBot1.SetConfig(10, "KRW-ETH", "minute")

	aBot2 := A_bot.NewBot()
	aBot2.SetConfig(10, "KRW-ETH", "minute", 200, 5)
	// 또는 aBot2.SetDefaultConfig()

	aBot1.Run()
	aBot2.Run()
}
