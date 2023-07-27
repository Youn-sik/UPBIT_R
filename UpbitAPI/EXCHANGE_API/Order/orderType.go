package Order

type Order struct{}

type OrderChanceInfo struct {
	BidFee      string   `json:"bid_fee"`       // 매수 수수료비율
	AskFee      string   `json:"ask_fee"`       // 매도 수수료 비율
	MakerBidFee string   `json:"maker_bid_fee"` // 매수 시 메이커 수수료 비율
	MakerAskFee string   `json:"maker_ask_fee"` // 매도 시 메이커 수수료 비율
	Market      struct { // 마켓에 대한 정보
		Id         string   `json:"id"`          // 마켓 ID
		Name       string   `json:"name"`        // 마켓 이름
		OrderTypes []string `json:"order_types"` // 지원 주문 방식 (만료)
		OrderSides []string `json:"order_sides"` // 지원 주문 종류
		BidTypes   []string `json:"bid_types"`   // 매수 주문 지원 방식
		AskTypes   []string `json:"ask_types"`   // 매도 주문 지원 방식
		Bid        struct { // 매수 시 제약사항
			Currency string `json:"currency"`  // 화폐를 의미하는 영문 대문자 코드
			MinTotal string `json:"min_total"` // 최소 매도/매수 금액
			// PriceUnit string `json:"price_unit"` // 주문금액 단위
		} `json:"bid"`
		Ask struct { // 매도 시 제약사항
			Currency string `json:"currency"`  // 화폐를 의미하는 영문 대문자 코드
			MinTotal string `json:"min_total"` // 최소 매도/매수 금액
			// PriceUnit string `json:"price_unit"` // 주문금액 단위
		} `json:"ask"`
		MaxTotal string `json:"max_total"` // 최대 매도/매수 금액
		State    string `json:"state"`     // 마켓 운영 상태
	} `json:"market"`
	BidAccount struct { // 매수 시 사용하는 화폐의 계좌 상태
		Currency            string `json:"currency"`               // 화폐를 의미하는 영문 대문자 코드
		Balance             string `json:"balance"`                // 주문가능 금액/수량
		Locked              string `json:"locked"`                 // 주문 중 묶여있는 금액/수량
		AvgBuyPrice         string `json:"avg_buy_price"`          // 매수 평균가
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"` // 매수 평균가 수정 여부
		UnitCurrency        string `json:"unit_currency"`          // 평단가 기준 화폐
	} `json:"bid_account"`
	AskAccount struct { // 매도 시 사용하는 화폐의 계좌 상태
		Currency            string `json:"currency"`               // 화폐를 의미하는 영문 대문자 코드
		Balance             string `json:"balance"`                // 주문가능 금액/수량
		Locked              string `json:"locked"`                 // 주문 중 묶여있는 금액/수량
		AvgBuyPrice         string `json:"avg_buy_price"`          // 매수 평균가
		AvgBuyPriceModified bool   `json:"avg_buy_price_modified"` // 매수 평균가 수정 여부
		UnitCurrency        string `json:"unit_currency"`          // 평단가 기준 화폐
	} `json:"ask_account"`
}

type OrderDetailInfo struct {
}

type OrderInfo struct {
}

type OrderCancelInfo struct {
}
