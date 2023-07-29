package B_bot

// B-bot 은 A-bot 1개당 1개가 배치 받아야한다.
// bot 에 따라 원하는 알고리즘 및 목적 수익량 등 다르게 설정이 들어간다.
// 채널로 통신하는 상태에서 봇의 종류를 알아차릴 수 있는 방법은 ?
// A-bot 쪽에서 B-bot 생성 매서드를 통해 생성하고
// A-bot 패키지 내에서 B-bot 객체 내부의 채널로 데이터 전송,
// B-bot 내부 채널의 리시버 함수는 B-bot 패키지 내에서 처리하도록 하면 ?

// Create B New Bot
func NewBot() B_botType {
	return B_botType{}
}

func (b *B_botType) Run() {
	b.run()
}
