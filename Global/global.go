package Global

import "time"

var UpClient UpbitClient

func ConvertTimestampToKST(ts int) string {
	seconds := ts / 1000
	kstTime := time.Unix(int64(seconds), 0).In(time.FixedZone("KST", 9*60*60))
	return kstTime.Format("2006-01-02T15:04:05")
}
