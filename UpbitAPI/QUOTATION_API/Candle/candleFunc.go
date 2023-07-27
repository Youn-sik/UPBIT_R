package Candle

import (
	"strconv"
)

func getMinCandleInputCheck(getMinCandleInputType GetMinCandleInputType) bool {
	var result bool
	gMCIT_MU := getMinCandleInputType.MinUnit
	gMCIT_T := getMinCandleInputType.To

	result = (gMCIT_MU == 0 ||
		(gMCIT_MU != 1 &&
			gMCIT_MU != 3 &&
			gMCIT_MU != 5 &&
			gMCIT_MU != 10 &&
			gMCIT_MU != 15 &&
			gMCIT_MU != 30 &&
			gMCIT_MU != 60 &&
			gMCIT_MU != 240) ||
		getMinCandleInputType.Count == 0 || getMinCandleInputType.Market == "")

	if gMCIT_T != "" {
		// 여기서 string regex 처리 필요
		result = (result || gMCIT_T == "")
	}

	return result
}

func setMinCandleInputToParameterType(getMinCandleInputType GetMinCandleInputType) GetMinCandleInputParameterType {
	var getMinCandleInputParameterType GetMinCandleInputParameterType

	getMinCandleInputParameterType.Count = strconv.Itoa(getMinCandleInputType.Count)
	getMinCandleInputParameterType.Market = getMinCandleInputType.Market
	getMinCandleInputParameterType.MinUnit = strconv.Itoa(getMinCandleInputType.MinUnit)
	getMinCandleInputParameterType.To = getMinCandleInputType.To

	return getMinCandleInputParameterType
}

func setDayCandleInputToParameterType(getDayCandleInpuyType GetDayCandleInputType) GetDayCandleInputParameterType {
	var getDayCandleInputParameterType GetDayCandleInputParameterType

	getDayCandleInputParameterType.Market = getDayCandleInpuyType.Market
	getDayCandleInputParameterType.To = getDayCandleInpuyType.To
	getDayCandleInputParameterType.Count = strconv.Itoa(getDayCandleInpuyType.Count)
	getDayCandleInputParameterType.ConvertingPriceUnit = getDayCandleInpuyType.ConvertingPriceUnit

	return getDayCandleInputParameterType
}

func setWeekMonthCandleInputToParameterType(getWeekCandleInputType GetWeekMonthCandleInputType) GetWeekMonthCandleInputParameterType {
	var getWeekMonthCandleInputParameterType GetWeekMonthCandleInputParameterType

	getWeekMonthCandleInputParameterType.Market = getWeekCandleInputType.Market
	getWeekMonthCandleInputParameterType.To = getWeekCandleInputType.To
	getWeekMonthCandleInputParameterType.Count = strconv.Itoa(getWeekCandleInputType.Count)

	return getWeekMonthCandleInputParameterType
}
