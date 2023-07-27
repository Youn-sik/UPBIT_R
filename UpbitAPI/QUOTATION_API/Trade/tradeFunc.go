package Trade

import "strconv"

func setTradeTickInputToParameterType(getTradeTickInputType GetTradeTickInputType) GetTradeTickInputParameterType {
	var getTradeTickInputParameterType GetTradeTickInputParameterType

	getTradeTickInputParameterType.Market = getTradeTickInputType.Market
	getTradeTickInputParameterType.To = getTradeTickInputType.To
	getTradeTickInputParameterType.Count = strconv.Itoa(getTradeTickInputType.Count)
	getTradeTickInputParameterType.Cursor = getTradeTickInputType.Cursor
	if getTradeTickInputType.DaysAgo == 0 {
		getTradeTickInputParameterType.DaysAgo = ""
	} else {
		getTradeTickInputParameterType.DaysAgo = strconv.Itoa(getTradeTickInputType.DaysAgo)
	}

	return getTradeTickInputParameterType
}
