package Order

import "strconv"

func setGetOrderQueryString(uuid, identifier string) string {
	var queryString string

	if uuid != "" && identifier != "" {
		queryString = "?uuid=" + uuid + "&identifier=" + identifier
	} else if uuid != "" && identifier == "" {
		queryString = "?uuid=" + uuid
	} else if uuid == "" && identifier != "" {
		queryString = "?identifier=" + identifier
	}

	return queryString
}

func setGetOrderListQueryString(
	market string,
	uuids, identifiers []string,
	state string,
	states []string,
	page, limit int,
	orderBy string,
) string {
	var queryString string

	queryString = "?market=" + market
	for _, val := range uuids {
		queryString += "&uuids=" + val
	}
	for _, val := range identifiers {
		queryString += "&identifiers=" + val
	}
	queryString += "&state=" + state
	for _, val := range states {
		queryString += "&states=" + val
	}

	if page != 0 {
		pageStr := strconv.Itoa(page)
		queryString += "&page=" + pageStr
	} else {
		queryString += "&page=1"
	}

	if limit != 0 {
		limitStr := strconv.Itoa(limit)
		queryString += "&limit=" + limitStr
	} else {
		queryString += "&limit=100"
	}

	if orderBy != "" {
		queryString += "&order_by=" + orderBy
	} else {
		queryString += "&order_by=desc"
	}

	return queryString
}

func setOrderCancelQueryString(uuid, identifier string) string {
	var queryString string

	queryString = "?uuid=" + uuid
	queryString += "&identifier=" + identifier

	return queryString
}
