package todo

import "time"

// TODO: move this function into helper package
func ParseTimeToEpoch(dateStr string) (int64, error) {
    parsedTime, err := time.Parse(time.DateOnly, dateStr)

    if err != nil {
        return -1, err
    }
    epochTimeMs := parsedTime.UnixNano() / int64(time.Millisecond)
    return epochTimeMs, nil
}
