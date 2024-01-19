package utils

import (
	"time"
)

func GetCurrentTimeInEpoch() int64 {
	now := time.Now()
	ms := now.UnixNano() / int64(time.Millisecond)
	return ms
}

func ParseDateToEpoch(dateStr string) (int64, error) {
	parsedTime, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return -1, err
	}
	epochTimeMs := parsedTime.UnixNano() / int64(time.Millisecond)
	return epochTimeMs, nil
}

func ParseEpochToDate(epochTimeMs int64) string {
	timestamp := time.Unix(0, epochTimeMs*int64(time.Millisecond))
	dateStr := timestamp.Format(time.DateOnly)
	return dateStr
}
