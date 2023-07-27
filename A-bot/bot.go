package A_bot

// Create A New Bot
func NewBot() A_botType {
	return A_botType{}
}

// Set Default Configuration About Bot
func (a *A_botType) SetDefaultConfig() {
	a.setDefaultConfig()
}

// Set Configuration With Given Data About Bot
func (a *A_botType) SetConfig(interval int, marketName, candleType string, candleCount, candleUnit int) {
	a.setConfig(interval, marketName, candleType, candleCount, candleUnit)
}

// Start Job
func (a *A_botType) Run() {
	a.run()
	//log.Println(a.CollectionInterval, a.MarketName, a.CandleType)
}

// A_bot 을 실행 시키면 제 역학을 진행하며
// 해당 역할을 위한 config 정보를 가져와 미리 세팅한다. -> 목표하는 마켓, 원하는 캔들, 현재 호가 수집 주기
// 봇은 각각 따로 생성하여 관리 및 config 할 수 있어야한다.
