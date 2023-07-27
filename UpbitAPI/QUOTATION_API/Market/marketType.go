package Market

type Market struct{}

type MarketInfo struct {
	Market      string `json:"market"`
	KoreanName  string `json:"korean_name"`
	EnglishName string `json:"english_name"`
}

/*
{"market":"BTC-GAL","korean_name":"갤럭시","english_name":"Galxe"}
{"market":"BTC-GAL","korean_name":"갤럭시","english_name":"Galxe"}
*/
