package A_bot

type A_botType struct {
	collectionInterval int
	marketName         string
	candleType         string // minute, day, week, month
	candleCount        int    // max 200
	candleUnit         int    // only minute candel -> 1,3,5,15,10,30,60,240
}
