package utils

import (
	"strconv"
	"time"
)

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func MergeMap(m1, m2 map[string]string) {
	for k, v := range m2 {
		m1[k] = v
	}
}

func MsToTime(ms int64) (time.Time, error) {
	str := Int64ToString(ms)
	msInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, msInt*int64(time.Millisecond)), nil
}
